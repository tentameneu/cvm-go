package main

import (
	"fmt"
)

func main() {
	cvm := &cvm{
		stream: generateRepeatingStream(100000000, 5000000),
		buffer: newTreapBuffer(10000),
	}
	n := cvm.run()
	cvm.buffer.printBasicInfo()
	fmt.Printf("Estimated distinct elements: %d\n", n)
}

func generateRepeatingStream(total, distinct int) []int {
	stream := make([]int, total)
	for i := 0; i < total; i++ {
		stream[i] = i % distinct
	}
	return stream
}
