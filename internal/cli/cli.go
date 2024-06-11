package cli

import (
	"flag"
	"fmt"

	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/cvm"
	"github.com/tentameneu/cvm-go/internal/stream"
)

var streamType = flag.String("stream-type", "incremental", "how to generate test stream of elements. valid values are: [incremental, random]")
var total = flag.Int("total", 100_000_000, "total number of elements in generated test stream")
var distinct = flag.Int("distinct", 5_000_000, "number of distincts elements in generated test stream")
var randomMin = flag.Int("random-min", 0, "used in random stream generator - generates values in range [random-min, random-max]")
var randomMax = flag.Int("random-max", 10_000_000, "used in random stream generator - generates values in range [random-min, random-max]")
var bufferSize = flag.Int("buffer-size", 10_000, "number of elements that can be stored in buffer while processing stream")

var generateConfigParams = func() map[string]any {
	return map[string]any{
		"streamType": *streamType,
		"total":      *total,
		"distinct":   *distinct,
		"randomMin":  *randomMin,
		"randomMax":  *randomMax,
		"bufferSize": *bufferSize,
	}
}

func Parse() (*cvm.CVMRunner, error) {
	flag.Usage = usage
	flag.Parse()
	return processArgs()
}

func processArgs() (*cvm.CVMRunner, error) {
	conf, err := config.NewConfig(generateConfigParams())
	if err != nil {
		return nil, err
	}

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
