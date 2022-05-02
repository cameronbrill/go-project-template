package cli

import (
	"flag"
	"net/http"
	"time"
)

const (
	DefaultIdleTimeout = 30 * time.Second
)

func Run(args []string) int {
	var app cli
	err := app.validateArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		panic(err)
	}
	return 0
}

type cli struct {
	hc   http.Client
	url  string
	mock bool
}

func (app *cli) validateArgs(args []string) error {
	app.hc = *http.DefaultClient
	fl := flag.NewFlagSet("example", flag.ContinueOnError)
	fl.StringVar(&app.url, "u", "https://example.com", "Request URL")
	fl.DurationVar(&app.hc.Timeout, "t", DefaultIdleTimeout, "Client timeout")
	fl.BoolVar(&app.mock, "mock", false, "Mock request")
	if err := fl.Parse(args); err != nil {
		return err
	}
	return nil
}

func (app *cli) run() error {
	var resp ExampleRes
	if app.mock {
		if err := fetchJSON(&app.hc, app.url, &resp); err != nil {
			return err
		}
	}
	return nil
}
