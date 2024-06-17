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
	logging.InitializeLogger(os.Stdout)

	err := cli.Parse()
	if err != nil {
		log().Error("Error while parsing CLI arguments", "err", err.Error())
		os.Exit(3)
	}

	if err := logging.InitializeLogger(os.Stdout); err != nil {
		log().Error("Error while initializing logger", "err", err.Error())
		os.Exit(4)
	}

	streamGenerator, err := stream.NewStreamGenerator()
	if err != nil {
		log().Error("Error while creating stream generator", "err", err.Error())
		os.Exit(5)
	}

	stream, err := streamGenerator.Generate()
	if err != nil {
		log().Error("Error while generating stream", "err", err.Error())
		os.Exit(6)
	}
	runner := cvm.NewCVMRunner(stream)
	runner.Run()
}
