package pipeline

import (
	"context"
	"strings"
	"testing"

	test "github.com/cameronbrill/go-project-template/test"
)

type stepTest[T any, U comparable] struct {
	name           string
	transform      func(T) (U, error)
	input          []T
	expectedOutput map[U]int
}

func (st stepTest[T, U]) Name() string {
	return st.name
}

func (st stepTest[T, U]) Test(t *testing.T) {
	ctx := context.Background()
	inChan := Producer(ctx, st.input)
	outChan := make(chan U, len(st.input))
	errChan := make(chan error)

	Step(context.Background(), inChan, outChan, errChan, st.transform)

	if len(st.input) != len(outChan) {
		t.Fatalf("expected %d items in outChan, got %d", len(st.input), len(outChan))
	}

	for i := 0; i < len(st.input); i++ {
		out := <-outChan
		if _, ok := st.expectedOutput[out]; !ok {
			t.Fatalf("expected %v in outChan, got %v", st.expectedOutput[out], out)
		}
		st.expectedOutput[out]--
		if st.expectedOutput[out] == 0 {
			delete(st.expectedOutput, out)
		}
	}
}

func TestStep(t *testing.T) {
	tests := []test.Testable{
		stepTest[string, string]{
			name:      "test string to string",
			transform: func(s string) (string, error) { return strings.ToLower(s), nil },
			input:     []string{"abc", "ABC", "DEF", "", "1234"},
			expectedOutput: map[string]int{
				"abc":  2,
				"def":  1,
				"":     1,
				"1234": 1,
			},
		},
		stepTest[string, int]{
			name:           "test string to int",
			transform:      func(s string) (int, error) { return len(s), nil },
			input:          []string{"FOO", "BAR", "BAX", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
			expectedOutput: map[int]int{3: 3, 26: 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name(), tc.Test)
	}
}
