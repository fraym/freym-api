package deployments

import (
	"context"

	"github.com/fraym/freym-api/go/proto/deployments/managementpb"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (c *MockClient) DeploySchema(
	ctx context.Context,
	schemaString string,
	options *DeploymentOptions,
) (*int64, error) {
	args := c.Called(ctx, schemaString, options)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*int64), args.Error(1)
}

func (c *MockClient) ActivateDeployment(ctx context.Context, deploymentId int64) error {
	return c.Called(ctx, deploymentId).Error(0)
}

func (c *MockClient) ConfirmDeployment(ctx context.Context, deploymentId int64) error {
	return c.Called(ctx, deploymentId).Error(0)
}

func (c *MockClient) RollbackDeployment(ctx context.Context, target managementpb.DeploymentTarget) error {
	return c.Called(ctx, target).Error(0)
}

func (c *MockClient) RollbackDeploymentById(ctx context.Context, deploymentId int64) error {
	return c.Called(ctx, deploymentId).Error(0)
}

func (c *MockClient) Close() error {
	return c.Called().Error(0)
}
