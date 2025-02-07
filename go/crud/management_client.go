package crud

import (
	"context"

	"github.com/fraym/freym-api/go/crud/config"
	"github.com/fraym/freym-api/go/crud/management"
	"github.com/fraym/freym-api/go/proto/crud/managementpb"
)

type MigrationStatus struct {
	Done bool
}

type ManagementClient interface {
	DeploySchema(
		ctx context.Context,
		namespace string,
		projectionTypes []ObjectType,
		crudTypes []ObjectType,
		nestedTypes []ObjectType,
		enums []EnumType,
		options *DeploymentOptions,
	) error
	ConfirmSchema(ctx context.Context, deploymentId int64) error
	RollbackSchema(ctx context.Context, deploymentId int64) error
	GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error)
	Close() error
}

type crudManagementClient struct {
	clientCloseFn func() error
	client        managementpb.ServiceClient
}

func NewManagementClient(conf *config.Config) (ManagementClient, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	client, clientCloseFn, err := management.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &crudManagementClient{
		clientCloseFn: clientCloseFn,
		client:        client,
	}, nil
}

func (c *crudManagementClient) DeploySchema(
	ctx context.Context,
	namespace string,
	projectionTypes []ObjectType,
	crudTypes []ObjectType,
	nestedTypes []ObjectType,
	enums []EnumType,
	options *DeploymentOptions,
) error {
	newProjectionTypes, err := objectTypeListToPb(projectionTypes)
	if err != nil {
		return err
	}

	newCrudTypes, err := objectTypeListToPb(crudTypes)
	if err != nil {
		return err
	}

	newNestedTypes, err := objectTypeListToPb(nestedTypes)
	if err != nil {
		return err
	}

	newEnums, err := enumTypeListToPb(enums)
	if err != nil {
		return err
	}

	_, err = c.client.DeploySchema(ctx, managementpb.DeploySchemaRequest_builder{
		Namespace:       namespace,
		ProjectionTypes: newProjectionTypes,
		CrudTypes:       newCrudTypes,
		NestedTypes:     newNestedTypes,
		EnumTypes:       newEnums,
		Options:         options.toPb(),
	}.Build())
	return err
}

func (c *crudManagementClient) ConfirmSchema(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ConfirmSchema(ctx, managementpb.ConfirmSchemaRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *crudManagementClient) RollbackSchema(ctx context.Context, deploymentId int64) error {
	_, err := c.client.RollbackSchema(ctx, managementpb.RollbackSchemaRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *crudManagementClient) GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error) {
	response, err := c.client.GetSchemaDeployment(ctx, managementpb.GetSchemaDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return response.GetProgress(), err
}

func (c *crudManagementClient) Close() error {
	return c.clientCloseFn()
}
