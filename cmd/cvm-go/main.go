package main

import (
	"fmt"

	"github.com/tentameneu/cvm-go/internal/cvm"
	"github.com/tentameneu/cvm-go/internal/stream"
)

func main() {
	generatorArgs := map[string]interface{}{
		"total":    100000000,
		"distinct": 5000000,
	}
	streamGenerator, _ := stream.NewStreamGenerator("repeating", generatorArgs)
	runner := cvm.NewCVMRunner(
		streamGenerator.Generate(),
		10000,
	)
	n := runner.Run()
	runner.PrintBufferBasicInfo()
	fmt.Printf("Estimated number of distinct elements: %d\n", n)
}
