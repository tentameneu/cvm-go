package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/tentameneu/cvm-go/internal/cvm"
	"github.com/tentameneu/cvm-go/internal/stream"
)

var genType = flag.String("gen-type", "repeating", "how to generate test stream of elements. valid values are: [repeating]")
var total = flag.Int("total", 100000000, "total number of elements in generated test stream")
var distinct = flag.Int("distinct", 5000000, "number of distincts elements in generated test stream")
var bufferSize = flag.Int("buffer-size", 10000, "number of elements that can be stored in buffer while processing stream")

func Parse() (*cvm.CVMRunner, error) {
	flag.Usage = usage
	flag.Parse()
	if *total < *distinct {
		printAndExit("Total number of elements can't be smaller than distinct number of elements!")
	}
	switch *genType {
	case "repeating":
		return createRepeatingStreamRunner()
	default:
		fmt.Fprintf(os.Stderr, "Unknown stream generator type '%s'\n\n", *genType)
		flag.PrintDefaults()
		os.Exit(1)
		return nil, nil
	}
}

func createRepeatingStreamRunner() (*cvm.CVMRunner, error) {
	generatorArgs := map[string]interface{}{
		"total":    *total,
		"distinct": *distinct,
	}
	streamGenerator, err := stream.NewStreamGenerator("repeating", generatorArgs)
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

func printAndExit(text string) {
	fmt.Fprintln(os.Stderr, text)
	os.Exit(2)
}
