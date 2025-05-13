package auth

import (
	"context"

	"github.com/fraym/freym-api/go/proto/auth/managementpb"
	"github.com/stretchr/testify/mock"
)

type MockManagementClient struct {
	mock.Mock
}

func (c *MockManagementClient) DeploySchema(
	ctx context.Context,
	deploymentId int64,
	namespace string,
	permissions []string,
	options *DeploymentOptions,
) error {
	return c.Called(ctx, deploymentId, namespace, permissions, options).Error(0)
}

func (c *MockManagementClient) ActivateSchema(ctx context.Context, deploymentId int64) error {
	return c.Called(ctx, deploymentId).Error(0)
}

func (c *MockManagementClient) ConfirmSchema(ctx context.Context, deploymentId int64) error {
	return c.Called(ctx, deploymentId).Error(0)
}

func (c *MockManagementClient) RollbackSchema(
	ctx context.Context,
	namespace string,
	target managementpb.DeploymentTarget,
) error {
	return c.Called(ctx, namespace, target).Error(0)
}

func (c *MockManagementClient) RollbackSchemaByDeployment(ctx context.Context, deploymentId int64) error {
	return c.Called(ctx, deploymentId).Error(0)
}

func (c *MockManagementClient) GetSchemaDeployment(ctx context.Context, deploymentId int64) (uint32, error) {
	args := c.Called(ctx, deploymentId)
	return args.Get(0).(uint32), args.Error(1)
}

func (c *MockManagementClient) Close() error {
	return c.Called().Error(0)
}
