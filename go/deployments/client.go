package deployments

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/fraym/freym-api/go/deployments/config"
	"github.com/fraym/freym-api/go/deployments/management"
	"github.com/fraym/freym-api/go/graphql/parser"
	"github.com/fraym/freym-api/go/graphql/schema"
	"github.com/fraym/freym-api/go/proto/deployments/managementpb"
	"github.com/fraym/golog"
	"github.com/samber/lo"
)

var ErrAlreadyClosed = fmt.Errorf("client already closed")

type DeploymentOptions struct {
	DangerouslyRemoveGdprFields bool
	SkipServices                []string
	Force                       bool
	Target                      managementpb.DeploymentTarget
}

type Client interface {
	DeploySchema(ctx context.Context, schemaString string, options *DeploymentOptions) (*int64, error)
	ActivateDeployment(ctx context.Context, deploymentId int64) error
	ConfirmDeployment(ctx context.Context, deploymentId int64) error
	RollbackDeployment(ctx context.Context) error
	Close() error
}

type deploymentsClient struct {
	clientCloseFn func() error
	client        managementpb.ServiceClient
	logger        golog.Logger
	conf          *config.Config

	mx                    *sync.Mutex
	isMigrating           bool
	migratingDeploymentId int64
	isClosed              bool
}

func NewClient(
	conf *config.Config,
	logger golog.Logger,
) (Client, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	client, clientCloseFn, err := management.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &deploymentsClient{
		clientCloseFn: clientCloseFn,
		client:        client,
		logger:        logger,
		conf:          conf,
		mx:            &sync.Mutex{},
		isMigrating:   false,
		isClosed:      false,
	}, nil
}

func (c *deploymentsClient) DeploySchema(ctx context.Context, schemaString string, options *DeploymentOptions) (*int64, error) {
	requestData, err := c.getDataFromString(schemaString)
	if err != nil {
		return nil, err
	}

	c.mx.Lock()
	if c.isClosed {
		return nil, ErrAlreadyClosed
	}

	c.isMigrating = true
	c.mx.Unlock()

	requestData.SetOptions(managementpb.DeploymentOptions_builder{
		DangerouslyRemoveGdprFields: options.DangerouslyRemoveGdprFields,
		SkipServices:                options.SkipServices,
		Force:                       options.Force,
		Target:                      options.Target,
	}.Build())

	response, err := c.client.CreateDeployment(ctx, requestData)
	if err != nil {
		c.mx.Lock()
		c.isMigrating = false
		c.mx.Unlock()
		return nil, err
	}

	errChan := make(chan error)
	finishDone := make(chan bool)

	go func(deploymentId int64) {
		for {
			status, err := c.client.GetDeployment(context.Background(), managementpb.GetDeploymentRequest_builder{
				DeploymentId: deploymentId,
			}.Build())
			if err != nil {
				errChan <- err
				return
			}

			if status == nil {
				continue
			}

			allDone := true

		componentLoop:
			for _, component := range status.GetComponents() {
				if component.GetProgress() != 100 {
					allDone = false
					break componentLoop
				}
			}

			if allDone {
				break
			}

			time.Sleep(time.Second)
		}

		finishDone <- true
	}(response.GetDeploymentId())

	select {
	case err := <-errChan:
		c.logger.Info().WithError(err).Write("rolling back migration due to error")

		if _, rollbackErr := c.client.RollbackDeployment(context.Background(), managementpb.RollbackDeploymentRequest_builder{
			DeploymentId: response.GetDeploymentId(),
			Namespace:    c.conf.Namespace,
		}.Build()); rollbackErr != nil {
			return nil, rollbackErr
		}

		c.mx.Lock()
		c.isMigrating = false
		c.migratingDeploymentId = 0
		c.mx.Unlock()

		return nil, err
	case <-finishDone:
		c.mx.Lock()
		c.isMigrating = false
		c.migratingDeploymentId = 0
		c.mx.Unlock()

		deploymentId := response.GetDeploymentId()
		return &deploymentId, nil
	}
}

func (c *deploymentsClient) ActivateDeployment(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ActivateDeployment(context.Background(), managementpb.ActivateDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *deploymentsClient) ConfirmDeployment(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ConfirmDeployment(context.Background(), managementpb.ConfirmDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *deploymentsClient) RollbackDeployment(ctx context.Context) error {
	_, err := c.client.RollbackDeployment(context.Background(), managementpb.RollbackDeploymentRequest_builder{
		Namespace: c.conf.Namespace,
	}.Build())
	return err
}

func (c *deploymentsClient) Close() error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.isClosed = true

	if c.isMigrating && c.migratingDeploymentId > 0 {
		c.logger.Info().Write("rolling back unfinished migration because of shutdown")

		request := managementpb.RollbackDeploymentRequest_builder{
			Namespace: c.conf.Namespace,
		}

		if c.migratingDeploymentId > 0 {
			request.DeploymentId = c.migratingDeploymentId
		}

		if _, err := c.client.RollbackDeployment(context.Background(), request.Build()); err != nil {
			return err
		}
	}

	return c.clientCloseFn()
}

func (c *deploymentsClient) getDataFromString(schemaString string) (*managementpb.CreateDeploymentRequest, error) {
	s, err := parser.GetSchemaFromString(schemaString)
	if err != nil {
		return nil, err
	}

	crudObjects := lo.Filter(s.Objects, func(obj schema.Object, _ int) bool {
		_, err := obj.GetDirective("crudType")
		return err == nil
	})

	projectionObjects := lo.Filter(s.Objects, func(obj schema.Object, _ int) bool {
		_, err := obj.GetDirective("upsertOn")
		_, err2 := obj.GetDirective("dangerouslyUpsertOn")
		return err == nil || err2 == nil
	})

	nestedObjects := lo.Filter(s.Objects, func(obj schema.Object, _ int) bool {
		if _, err := obj.GetDirective("crudType"); err == nil {
			return false
		}

		if _, err := obj.GetDirective("upsertOn"); err == nil {
			return false
		}

		if _, err := obj.GetDirective("dangerouslyUpsertOn"); err == nil {
			return false
		}

		return true
	})

	crudTypes, err := objectsToPb(crudObjects)
	if err != nil {
		return nil, err
	}

	projectionTypes, err := objectsToPb(projectionObjects)
	if err != nil {
		return nil, err
	}

	nestedTypes, err := objectsToPb(nestedObjects)
	if err != nil {
		return nil, err
	}

	enumTypes, permissions, err := enumsAndPermissionsToPb(s.Enums)
	if err != nil {
		return nil, err
	}

	return managementpb.CreateDeploymentRequest_builder{
		CrudTypes:       crudTypes,
		ProjectionTypes: projectionTypes,
		EnumTypes:       enumTypes,
		NestedTypes:     nestedTypes,
		Permissions:     permissions,
		Namespace:       c.conf.Namespace,
	}.Build(), nil
}
