package cvm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestStreamRunner(tota, distinct, bufferSize int) *CVM {
	return NewCVMRunner(newTestIncrementalStream(tota, distinct), bufferSize)
}

func TestRun(t *testing.T) {
	t.Run("SmallerBuffer", func(t *testing.T) {
		runner := newTestStreamRunner(1_000_000, 10_000, 1_000)
		n := runner.Run()
		assert.InDelta(t, 10_000, n, 1_000)
	})

	t.Run("ExactBuffer", func(t *testing.T) {
		runner := newTestStreamRunner(1_000_000, 10_000, 10_000)
		n := runner.Run()
		assert.Exactly(t, 10_000, n)
	})
}
