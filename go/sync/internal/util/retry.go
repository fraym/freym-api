package util

import "time"

func Retry(fn func() error, pause time.Duration, maxRetries int) error {
	if err := fn(); err == nil {
		return nil
	}

	var err error

	for i := 0; i < maxRetries; i++ {
		time.Sleep(pause)

		err = fn()

		if err == nil {
			return nil
		}
	}

	return err
}
