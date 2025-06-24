package projections

import (
	"context"

	"github.com/fraym/freym-api/go/projections/config"
	"github.com/fraym/freym-api/go/projections/management"
	"github.com/fraym/freym-api/go/proto/projections/managementpb"
)

type ManagementClient interface {
	DeploySchema(
		ctx context.Context,
		deploymentId int64,
		namespace string,
		projectionTypes []ObjectType,
		crudTypes []ObjectType,
		nestedTypes []ObjectType,
		enums []EnumType,
		views []View,
		baseViews []View,
		options *DeploymentOptions,
	) error
	ActivateSchema(ctx context.Context, deploymentId int64) error
	ConfirmSchema(ctx context.Context, deploymentId int64) error
	RollbackSchema(ctx context.Context, namespace string, target managementpb.DeploymentTarget) error
	RollbackSchemaByDeployment(ctx context.Context, deploymentId int64) error
	GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error)
	Close() error
}

type projectionsManagementClient struct {
	clientCloseFn func() error
	client        managementpb.ServiceClient
}

func NewManagementClient(
	conf *config.Config,
) (ManagementClient, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	client, clientCloseFn, err := management.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &projectionsManagementClient{
		clientCloseFn: clientCloseFn,
		client:        client,
	}, nil
}

func (c *projectionsManagementClient) DeploySchema(
	ctx context.Context,
	deploymentId int64,
	namespace string,
	projectionTypes []ObjectType,
	crudTypes []ObjectType,
	nestedTypes []ObjectType,
	enums []EnumType,
	views []View,
	baseViews []View,
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

	newViews, err := viewListToPb(views)
	if err != nil {
		return err
	}

	newBaseViews, err := viewListToPb(baseViews)
	if err != nil {
		return err
	}

	_, err = c.client.DeploySchema(ctx, managementpb.DeploySchemaRequest_builder{
		DeploymentId:    deploymentId,
		Namespace:       namespace,
		ProjectionTypes: newProjectionTypes,
		CrudTypes:       newCrudTypes,
		NestedTypes:     newNestedTypes,
		EnumTypes:       newEnums,
		Views:           newViews,
		BaseViews:       newBaseViews,
		Options:         options.toPb(),
	}.Build())
	return err
}

func (c *projectionsManagementClient) ActivateSchema(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ActivateSchema(ctx, managementpb.ActivateSchemaRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *projectionsManagementClient) ConfirmSchema(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ConfirmSchema(ctx, managementpb.ConfirmSchemaRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *projectionsManagementClient) RollbackSchema(
	ctx context.Context,
	namespace string,
	target managementpb.DeploymentTarget,
) error {
	_, err := c.client.RollbackSchema(ctx, managementpb.RollbackSchemaRequest_builder{
		Target:    target,
		Namespace: namespace,
	}.Build())
	return err
}

func (c *projectionsManagementClient) RollbackSchemaByDeployment(ctx context.Context, deploymentId int64) error {
	_, err := c.client.RollbackSchemaByDeployment(ctx, managementpb.RollbackSchemaByDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *projectionsManagementClient) GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error) {
	response, err := c.client.GetSchemaDeployment(ctx, managementpb.GetSchemaDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return response.GetProgress(), err
}

func (c *projectionsManagementClient) Close() error {
	return c.clientCloseFn()
}
