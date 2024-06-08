package main

import (
	"fmt"

	"github.com/tentameneu/cvm-go/internal/cvm"
)

func main() {
	runner := cvm.NewCVMRunner(
		cvm.GenerateRepeatingStream(100000000, 5000000),
		10000,
	)
	n := runner.Run()
	runner.PrintBufferBasicInfo()
	fmt.Printf("Estimated distinct elements: %d\n", n)
}
