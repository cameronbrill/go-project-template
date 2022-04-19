package pipeline

import (
	"context"
)

func Producer[T any](ctx context.Context, input []T) <-chan T {
	outChannel := make(chan T)

	go func() {
		defer close(outChannel)
		if err := ctx.Err(); err != nil {
			return
		}

		for _, v := range input {
			select {
			case <-ctx.Done():
				return
			case outChannel <- v:
			}
		}
	}()

	return outChannel
}
