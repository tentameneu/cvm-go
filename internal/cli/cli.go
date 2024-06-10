package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/cvm"
	"github.com/tentameneu/cvm-go/internal/stream"
)

var genType = flag.String("gen-type", "repeating", "how to generate test stream of elements. valid values are: [repeating]")
var total = flag.Int("total", 100000000, "total number of elements in generated test stream")
var distinct = flag.Int("distinct", 5000000, "number of distincts elements in generated test stream")
var bufferSize = flag.Int("buffer-size", 10000, "number of elements that can be stored in buffer while processing stream")

var generateConfigParams = func() map[string]any {
	return map[string]any{
		"genType":    *genType,
		"total":      *total,
		"distinct":   *distinct,
		"bufferSize": *bufferSize,
	}
}

func Parse() (*cvm.CVMRunner, error) {
	flag.Usage = usage
	flag.Parse()
	return processArgs()
}

func processArgs() (*cvm.CVMRunner, error) {
	switch *genType {
	case "repeating":
		conf, err := config.NewConfig(generateConfigParams())
		if err != nil {
			return nil, err
		}
		return createRepeatingStreamRunner(conf)
	default:
		fmt.Fprintf(os.Stderr, "Unknown stream generator type '%s'\n\n", *genType)
		flag.PrintDefaults()
		os.Exit(1)
		return nil, nil
	}
}

func createRepeatingStreamRunner(conf *config.Config) (*cvm.CVMRunner, error) {
	streamGenerator, err := stream.NewStreamGenerator(conf)
	if err != nil {
		return nil, err
	}
	return cvm.NewCVMRunner(streamGenerator.Generate(), *bufferSize), nil
}

func usage() {
	fmt.Print(`
This program runs CVM algorithm simulator.

Algorithm estimates number of distinct elments in stream of elements much bigger than available buffer size.

Usage:

cvm-go [arguments]

Supported arguments:

`)
	flag.PrintDefaults()
}
