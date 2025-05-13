package auth

import (
	"context"

	"github.com/fraym/freym-api/go/auth/config"
	"github.com/fraym/freym-api/go/auth/management"
	"github.com/fraym/freym-api/go/proto/auth/managementpb"
)

type DeploymentOptions struct {
	Target managementpb.DeploymentTarget
	Force  bool
}

func (o *DeploymentOptions) toPb() *managementpb.DeploymentOptions {
	if o == nil {
		return nil
	}

	return managementpb.DeploymentOptions_builder{
		Target: o.Target,
		Force:  o.Force,
	}.Build()
}

type ManagementClient interface {
	DeploySchema(
		ctx context.Context,
		deploymentId int64,
		namespace string,
		permissions []string,
		options *DeploymentOptions,
	) error
	ActivateSchema(ctx context.Context, deploymentId int64) error
	ConfirmSchema(ctx context.Context, deploymentId int64) error
	RollbackSchema(ctx context.Context, namespace string, target managementpb.DeploymentTarget) error
	RollbackSchemaByDeployment(ctx context.Context, deploymentId int64) error
	GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error)
	Close() error
}

type authManagementClient struct {
	clientCloseFn func() error
	client        managementpb.ServiceClient
}

func NewManagementClient(
	ctx context.Context,
	conf *config.Config,
) (ManagementClient, error) {
	if conf == nil {
		conf = config.NewDefaultConfig()
	}

	client, clientCloseFn, err := management.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &authManagementClient{
		clientCloseFn: clientCloseFn,
		client:        client,
	}, nil
}

func (c *authManagementClient) DeploySchema(
	ctx context.Context,
	deploymentId int64,
	namespace string,
	permissions []string,
	options *DeploymentOptions,
) error {
	_, err := c.client.DeploySchema(ctx, managementpb.DeploySchemaRequest_builder{
		DeploymentId: deploymentId,
		Namespace:    namespace,
		Permissions:  permissions,
		Options:      options.toPb(),
	}.Build())
	return err
}

func (c *authManagementClient) ActivateSchema(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ActivateSchema(ctx, managementpb.ActivateSchemaRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *authManagementClient) ConfirmSchema(ctx context.Context, deploymentId int64) error {
	_, err := c.client.ConfirmSchema(ctx, managementpb.ConfirmSchemaRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *authManagementClient) RollbackSchema(
	ctx context.Context,
	namespace string,
	target managementpb.DeploymentTarget,
) error {
	_, err := c.client.RollbackSchema(ctx, managementpb.RollbackSchemaRequest_builder{
		Namespace: namespace,
		Target:    target,
	}.Build())
	return err
}

func (c *authManagementClient) RollbackSchemaByDeployment(ctx context.Context, deploymentId int64) error {
	_, err := c.client.RollbackSchemaByDeployment(ctx, managementpb.RollbackSchemaByDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return err
}

func (c *authManagementClient) GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error) {
	response, err := c.client.GetSchemaDeployment(ctx, managementpb.GetSchemaDeploymentRequest_builder{
		DeploymentId: deploymentId,
	}.Build())
	return response.GetProgress(), err
}

func (c *authManagementClient) Close() error {
	return c.clientCloseFn()
}
