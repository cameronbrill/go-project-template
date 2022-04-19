package pipeline

import (
	"context"
	"testing"
)

func TestConsumer(t *testing.T) {
	input := []string{"FOO", "BAR", "BAX", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	ctx, cancel := context.WithCancel(context.Background())
	inChan := Producer(ctx, input)

	errChan := make(chan error)
	defer close(errChan)
	err := Consumer(ctx, cancel, inChan, errChan)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("canceled context", func(t *testing.T) {
		three := Producer(ctx, input)
		ctx, cancel = context.WithCancel(context.Background())
		cancel()

		err = Consumer(ctx, cancel, three, errChan)
		if err != context.Canceled {
			t.Fatal(err)
		}
	})
}
