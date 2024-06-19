package cvm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	t.Run("SmallerBuffer", func(t *testing.T) {
		runner := NewCVM(1_000, intTestComparator)
		var n int
		for _, element := range newTestIntStream(1_000_000, 10_000) {
			n = runner.Process(element)
		}
		assert.InDelta(t, 10_000, n, 1_000)
	})

	t.Run("ExactBuffer", func(t *testing.T) {
		runner := NewCVM(10_000, intTestComparator)
		var n int
		for _, element := range newTestIntStream(1_000_000, 10_000) {
			n = runner.Process(element)
		}
		assert.Exactly(t, 10_000, n)
	})
}
