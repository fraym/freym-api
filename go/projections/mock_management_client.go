package projections

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockManagementClient struct {
	mock.Mock
}

func (c *MockManagementClient) RegisterMigration(
	ctx context.Context,
	namespace string,
	dangerouslyRemoveGdprFields bool,
	projectionTypes []ObjectType,
	crudTypes []ObjectType,
	nestedTypes []ObjectType,
	enums []EnumType,
	views []View,
) error {
	return c.Called(ctx, namespace, dangerouslyRemoveGdprFields, projectionTypes, crudTypes, nestedTypes, enums, views).Error(0)
}

func (c *MockManagementClient) ApplyMigration(ctx context.Context, namespace string) error {
	return c.Called(ctx, namespace).Error(0)
}

func (c *MockManagementClient) CleanupMigration(ctx context.Context, namespace string, status map[string]int64) error {
	return c.Called(ctx, namespace, status).Error(0)
}

func (c *MockManagementClient) RollbackMigration(ctx context.Context, namespace string) error {
	return c.Called(ctx, namespace).Error(0)
}

func (c *MockManagementClient) GetMigrationStatus(ctx context.Context, namespace string) (*MigrationStatus, error) {
	args := c.Called(ctx, namespace)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*MigrationStatus), args.Error(1)
}

func (c *MockManagementClient) Close() error {
	return c.Called().Error(0)
}
