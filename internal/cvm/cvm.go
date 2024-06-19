package cvm

import (
	"fmt"
	"math/rand"

	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/logging"
)

var (
	log = logging.Logger
)

type CVMRunner struct {
	stream   []int
	buffer   *treapBuffer
	distinct int
}

func NewCVMRunner(stream []int) *CVMRunner {
	return &CVMRunner{
		stream:   stream,
		buffer:   newTreapBuffer(config.BufferSize(), func(x, y int) int { return x - y }),
		distinct: config.Distinct(),
	}
}

func (runner *CVMRunner) Run() int {
	log().Info("Starting CVM Algorithm...")

	p, m := 1.0, len(runner.stream)
	for i, a := range runner.stream {
		u := rand.Float64()
		log().Debug(fmt.Sprintf("Starting loop %d/%d", i+1, m), "p", p, "u", u, "Root", runner.rootString())
		logging.LogDeep("Deleting next element from buffer", "a", a)
		runner.buffer.delete(a)

		if u >= p {
			continue
		}
		if runner.buffer.currentSize < runner.buffer.maxSize {
			logging.LogDeep("|B| < s. Inserting new node", "a", a, "u", u)
			runner.buffer.insert(newNode(a, u))
			continue
		}
		if u > runner.buffer.root.priority {
			logging.LogDeep("Setting new 'p' by 'u'", "p", p, "u", u)
			p = u
		} else {
			logging.LogDeep("Setting new 'p' by max 'u' (root node priority)", "p", p, "u", u)
			p = runner.buffer.root.priority
			logging.LogDeep("Deleting root", "root.value", runner.buffer.root.value)
			runner.buffer.delete(runner.buffer.root.value)
			logging.LogDeep("Inserting new node", "a", a, "u", u)
			runner.buffer.insert(newNode(a, u))
		}
	}
	n := int(float64(runner.buffer.currentSize) / p)
	precision := float64(n) / float64(runner.distinct)
	if precision > 1.0 {
		precision = 1 / precision
	}

	log().Info("Done estimating number of distinct elements.", "N", n, "precision", fmt.Sprintf("%.2f%%", precision*100))
	log().Info(
		"Buffer status:",
		"Size", runner.buffer.GetCurrentSize(),
		"Root", runner.rootString(),
	)

	return n
}

func nodeString(node *node) string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("<Value: %d, Priority: %f>", node.value, node.priority)
}

func (runner *CVMRunner) rootString() string {
	return nodeString(runner.buffer.GetRoot())
}
