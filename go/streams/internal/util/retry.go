package util

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Retry(fn func() error, pause time.Duration, maxRetries int) error {
	err := fn()
	if err == nil {
		return nil
	}
	if errors.Is(err, context.Canceled) {
		return err
	}
	if status.Code(err) == codes.Unknown {
		return err
	}

	for i := 0; i < maxRetries; i++ {
		time.Sleep(pause)

		err = fn()
		if status.Code(err) == codes.Unknown {
			return err
		}
		if errors.Is(err, context.Canceled) {
			return err
		}
		if err == nil {
			return nil
		}
	}

	return err
}

func RetryWithResult[T any](fn func() (T, error), pause time.Duration, maxRetries int) (T, error) {
	var (
		err    error
		result T
	)

	result, err = fn()
	if err == nil {
		return result, nil
	}
	if status.Code(err) == codes.Unknown {
		return result, err
	}

	for i := 0; i < maxRetries; i++ {
		time.Sleep(pause)

		result, err = fn()
		if status.Code(err) == codes.Unknown {
			return result, err
		}
		if err == nil {
			return result, nil
		}
	}

	return result, err
}
