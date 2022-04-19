package pipeline

import (
	"context"
	"log"
)

func Consumer[T any](ctx context.Context, cancelFunc context.CancelFunc, values <-chan T, errors <-chan error) {
	for {
		select {
		case <-ctx.Done():
			log.Print(ctx.Err().Error())
			return
		case err := <-errors:
			if err != nil {
				log.Println("error: ", err.Error())
				cancelFunc()
			}
		case val, ok := <-values:
			if !ok {
				log.Print("done")
				return
			}
			log.Printf("Consumed: %v", val)
		}
	}
}
