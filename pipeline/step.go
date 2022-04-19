package pipeline

import (
	"context"
	"runtime"

	"golang.org/x/sync/semaphore"
)

func Step[In any, Out any](ctx context.Context,
	inChan <-chan In,
	outChan chan<- Out,
	errChan chan<- error,
	transform func(In) (Out, error),
) {
	defer close(outChan)

	limit := runtime.GOMAXPROCS(0)
	sem := semaphore.NewWeighted(int64(limit))

	for in := range inChan {
		select {
		case <-ctx.Done():
			print("fail")
			break
		default:
		}

		if err := sem.Acquire(ctx, 1); err != nil {
			print("lame")
			break
		}

		go func(in In) {
			defer sem.Release(1)
			res, err := transform(in)
			if err != nil {
				errChan <- err
				return
			}
			outChan <- res
		}(in)
	}

	if err := sem.Acquire(ctx, int64(limit)); err != nil {
		print("failed to acquire semaphore: ", err)
	}
}
