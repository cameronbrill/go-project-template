package main

import (
	"os"

	"github.com/cameronbrill/go-project-template/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
