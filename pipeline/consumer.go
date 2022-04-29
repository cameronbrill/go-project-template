package pipeline

import (
	"context"
	"log"
)

func Consumer[T any](ctx context.Context, cancelFunc context.CancelFunc, values <-chan T, errors <-chan error) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, ok := <-errors:
			if ok {
				log.Printf("error: %v", err)
				cancelFunc()
				return err
			}
		case val, ok := <-values:
			if !ok {
				return nil
			}
			log.Printf("Consumed: %v", val)
		default:
		}
	}
}
