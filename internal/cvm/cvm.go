package cvm

import (
	"io"
	"math/rand"
)

type CVMRunner struct {
	stream []int
	buffer *treapBuffer
}

func NewCVMRunner(stream []int, bufferSize int) *CVMRunner {
	return &CVMRunner{
		stream: stream,
		buffer: newTreapBuffer(bufferSize),
	}
}

func (runner *CVMRunner) PrintBufferBasicInfo(writer io.Writer) {
	runner.buffer.printBasicInfo(writer)
}

func (runner *CVMRunner) Run() int {
	p := 1.0
	for _, a := range runner.stream {
		runner.buffer.delete(a)
		u := rand.Float64()
		if u >= p {
			continue
		}
		if runner.buffer.currentSize < runner.buffer.maxSize {
			runner.buffer.insert(newNode(a, u))
			continue
		}
		if u > runner.buffer.root.priority {
			p = u
		} else {
			p = runner.buffer.root.priority
			runner.buffer.delete(runner.buffer.root.value)
			runner.buffer.insert(newNode(a, u))
		}
	}
	return int(float64(runner.buffer.currentSize) / p)
}
