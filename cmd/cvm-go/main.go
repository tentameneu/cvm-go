package main

import (
	"os"

	"github.com/tentameneu/cvm-go/internal/cli"
	"github.com/tentameneu/cvm-go/internal/logging"
)

var log = logging.Logger

func main() {
	logging.InitializeLogger(os.Stdout)

	runner, err := cli.Parse()
	if err != nil {
		log().Error("Error while parsing CLI arguments", "err", err.Error())
		os.Exit(3)
	}

	runner.Run()
}
