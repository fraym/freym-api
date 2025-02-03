package subscription

import (
	"fmt"

	"github.com/fraym/freym-api/go/proto/streams/managementpb"
)

var (
	ErrUnimplementedFunction = newError("unimplemented function")
	ErrUnknownGrpcRequest    = newError("unknown gRPC request")
)

func newError(msg string) error {
	return fmt.Errorf("%s: %s", "streams client", msg)
}

func defaultOnSubscribedFn(*managementpb.Subscribed) error {
	return fmt.Errorf("onSubscribedFn: %w", ErrUnimplementedFunction)
}

func defaultOnPanicFn(*managementpb.Panic) error {
	return fmt.Errorf("onPanicFn: %w", ErrUnimplementedFunction)
}

func defaultOnEventFn(*managementpb.Event, Connection) error {
	return fmt.Errorf("onEventFn: %w", ErrUnimplementedFunction)
}
