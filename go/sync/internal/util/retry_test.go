package util_test

import (
	"fmt"
	"testing"

	"github.com/fraym/freym-api/go/sync/internal/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testCaller struct {
	mock.Mock
}

func (c *testCaller) Call() error {
	return c.Called().Error(0)
}

func TestRetry(t *testing.T) {
	caller := &testCaller{}

	caller.On("Call").Return(fmt.Errorf("error")).Times(6)
	assert.Error(t, util.Retry(caller.Call, 0, 5))
	caller.AssertExpectations(t)

	caller.On("Call").Return(fmt.Errorf("error")).Times(5)
	caller.On("Call").Return(nil).Once()
	assert.NoError(t, util.Retry(caller.Call, 0, 5))
	caller.AssertExpectations(t)

	caller.On("Call").Return(fmt.Errorf("error")).Once()
	caller.On("Call").Return(nil).Once()
	assert.NoError(t, util.Retry(caller.Call, 0, 5))
	caller.AssertExpectations(t)

	caller.On("Call").Return(nil).Once()
	assert.NoError(t, util.Retry(caller.Call, 0, 5))
	caller.AssertExpectations(t)
}
