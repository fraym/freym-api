package util

import (
	"context"
)

func Iterate[T any](
	ctx context.Context,
	handle func(ctx context.Context, data T) error,
	loadData func(ctx context.Context, page int, perPage int) ([]T, error),
	stopLoadingMore func(ctx context.Context, data T) bool,
	perPage int,
	queueSize int,
) error {
	var readErr error
	queue := make(chan T, queueSize)
	readCtx, cancelRead := context.WithCancel(ctx)
	defer cancelRead()

	go func(ctx context.Context) {
		page := 0

		for {
			dataSlice, err := loadData(ctx, page, perPage)
			if err != nil {
				readErr = err
				close(queue)
				return
			}

			page++

			var last T

			for _, data := range dataSlice {
				queue <- data
				last = data
			}

			if stopLoadingMore(ctx, last) {
				close(queue)
				return
			}
		}
	}(readCtx)

	for data := range queue {
		if err := handle(ctx, data); err != nil {
			return err
		}
	}

	return readErr
}
