package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cameronbrill/go-project-template/pipeline"
)

func main() {
	source := []string{"FOO", "BAR", "BAX", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	readStream, err := pipeline.Producer(ctx, source)
	if err != nil {
		log.Fatal(err)
	}

	lowerStage := make(chan string)
	errorChannel := make(chan error)

	transformA := func(s string) (string, error) {
		return strings.ToLower(s), nil
	}

	go func() {
		pipeline.Step(ctx, readStream, lowerStage, errorChannel, transformA)
	}()

	type result struct {
		v string
		l int
	}

	titleStage := make(chan result)
	transformB := func(s string) (result, error) {
		if len(s) > 14 {
			return result{
				v: s,
				l: len(s),
			}, fmt.Errorf("invalid input string: %s", s)
		}

		return result{
			v: strings.Title(s),
			l: len(s),
		}, nil
	}

	go func() {
		pipeline.Step(ctx, lowerStage, titleStage, errorChannel, transformB)
	}()

	pipeline.Consumer(ctx, cancel, titleStage, errorChannel)
}
