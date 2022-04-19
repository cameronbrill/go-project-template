package pipeline

import (
	"context"
	"testing"
)

func TestProducer(t *testing.T) {
	input := []string{"FOO", "BAR", "BAX", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}
	ctx, cancel := context.WithCancel(context.Background())
	inChan := Producer(ctx, input)
	chanLen := 0
	for range inChan {
		chanLen++
	}

	if len(input) != chanLen {
		t.Fatalf("expected %d items in inChan, got %d", len(input), chanLen)
	}

	cancel()
	inChan = Producer(ctx, input)
	chanLen = 0
	for range inChan {
		chanLen++
	}
	if chanLen != 0 {
		t.Fatalf("expected %d items in inChan, got %d", 0, chanLen)
	}
}
