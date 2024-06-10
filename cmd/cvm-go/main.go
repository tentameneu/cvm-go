package main

import (
	"fmt"
	"os"

	"github.com/tentameneu/cvm-go/internal/cli"
)

func main() {
	runner, err := cli.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
	n := runner.Run()
	runner.PrintBufferBasicInfo(os.Stdout)
	fmt.Printf("Estimated number of distinct elements: %d\n", n)
}
