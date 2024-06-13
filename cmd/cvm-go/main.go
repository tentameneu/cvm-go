package main

import (
	"os"

	"github.com/tentameneu/cvm-go/internal/cli"
	"github.com/tentameneu/cvm-go/internal/cvm"
	"github.com/tentameneu/cvm-go/internal/logging"
	"github.com/tentameneu/cvm-go/internal/stream"
)

var log = logging.Logger

func main() {
	logging.InitializeLogger(os.Stdout, nil)

	conf, err := cli.Parse()
	if err != nil {
		log().Error("Error while parsing CLI arguments", "err", err.Error())
		os.Exit(3)
	}

	logging.InitializeLogger(os.Stdout, conf)

	streamGenerator, err := stream.NewStreamGenerator(conf)
	if err != nil {
		log().Error("Error while generating stream", "err", err.Error())
		os.Exit(4)
	}

	runner := cvm.NewCVMRunner(streamGenerator.Generate(), conf)
	runner.Run()
}
