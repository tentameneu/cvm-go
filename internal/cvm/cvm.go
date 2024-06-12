package cvm

import (
	"fmt"
	"math/rand"

	"github.com/tentameneu/cvm-go/internal/logging"
)

var log = logging.Logger

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

func (runner *CVMRunner) Run() int {
	log().Info("Starting CVM Algorithm...")

	p, m := 1.0, len(runner.stream)
	for i, a := range runner.stream {
		u := rand.Float64()
		log().Debug(fmt.Sprintf("Starting loop %d/%d", i+1, m), "p", p, "u", u, "Root", runner.rootString())
		runner.buffer.delete(a)

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
	n := int(float64(runner.buffer.currentSize) / p)

	log().Info("Done estimating number of distinct elements.", "N", n)
	log().Info(
		"Buffer status:",
		"Size", runner.buffer.GetCurrentSize(),
		"Root", runner.rootString(),
	)

	return n
}

func (runner *CVMRunner) rootString() string {
	return nodeString(runner.buffer.GetRoot())
}
