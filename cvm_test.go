package cvm

import (
	"math/rand"
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

// Test cases from original paper found at https://cs.stanford.edu/~knuth/papers/cvm-note.pdf
// Tests use different buffer size for each test stream. Every stream containing total of 1_000_000 elements.
// NOTE: These tests can fail due to randomness in algorithm specially for smaller buffer sizes. Try to run them multiple time.
func TestPaperStreams(t *testing.T) {
	// Random test stream. it contains 50_000 distinct elements generated randomly (value between 1 and 1_000_000), repeated 20 times in the same order.
	// Slightly different from original paper.
	t.Run("Random", func(t *testing.T) {
		newRandomTestStream := func() ([]int, int) {
			total := 1_000_000
			stream := make([]int, total)
			generated := make(map[int]bool)

			for i := 0; i < total; i++ {
				n := rand.Intn(9_000_000) + 1_000_000
				stream[i] = n
				generated[n] = true
			}
			return stream, len(generated)
		}

		t.Run("BufferSize=10", func(t *testing.T) {
			runner := NewCVM(10, intTestComparator)
			stream, n := newRandomTestStream()
			for _, element := range stream {
				runner.Process(element)
			}
			assert.InDelta(t, n, runner.N(), 0.5*float64(n))
		})

		t.Run("BufferSize=100", func(t *testing.T) {
			runner := NewCVM(100, intTestComparator)
			stream, n := newRandomTestStream()
			for _, element := range stream {
				runner.Process(element)
			}
			assert.InDelta(t, n, runner.N(), 0.2*float64(n))
		})

		t.Run("BufferSize=1_000", func(t *testing.T) {
			runner := NewCVM(1_000, intTestComparator)
			stream, n := newRandomTestStream()
			for _, element := range stream {
				runner.Process(element)
			}
			assert.InDelta(t, n, runner.N(), 0.1*float64(n))
		})

		t.Run("BufferSize=10_000", func(t *testing.T) {
			runner := NewCVM(10_000, intTestComparator)
			stream, n := newRandomTestStream()
			for _, element := range stream {
				runner.Process(element)
			}
			assert.InDelta(t, n, runner.N(), 0.03*float64(n))
		})

		t.Run("BufferSize=100_000", func(t *testing.T) {
			runner := NewCVM(100_000, intTestComparator)
			stream, n := newRandomTestStream()
			for _, element := range stream {
				runner.Process(element)
			}
			assert.InDelta(t, n, runner.N(), 0.01*float64(n))
		})
	})

	// Incremental test stream. It consists of 50_000 distinct elements generated incrementaly (0, 1, 2...49_999), repeated 20 times in the same order.
	t.Run("Incremental", func(t *testing.T) {
		t.Run("BufferSize=10", func(t *testing.T) {
			runner := NewCVM(10, intTestComparator)
			for _, element := range newTestIntStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 40_000)
		})

		t.Run("BufferSize=100", func(t *testing.T) {
			runner := NewCVM(100, intTestComparator)
			for _, element := range newTestIntStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 20_000)
		})

		t.Run("BufferSize=1_000", func(t *testing.T) {
			runner := NewCVM(1_000, intTestComparator)
			for _, element := range newTestIntStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 5_000)
		})

		t.Run("BufferSize=10_000", func(t *testing.T) {
			runner := NewCVM(10_000, intTestComparator)
			for _, element := range newTestIntStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 1_000)
		})

		t.Run("BufferSize=100_000", func(t *testing.T) {
			runner := NewCVM(100_000, intTestComparator)
			for _, element := range newTestIntStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.Equal(t, 50_000, runner.N())
		})
	})

	// Dual incremental test stream. It consists of 50_000 distinct elements generated incrementaly but repeated 20 times before incrementing (0, 0...1, 1...49_999, 49_999...).
	t.Run("DualIncremental", func(t *testing.T) {
		newIncrementalDualTestStream := func(total, distinct int) []int {
			repetitions := total / distinct
			stream := make([]int, total)
			for i := 0; i < distinct; i++ {
				for j := 0; j < repetitions; j++ {
					stream[i*repetitions+j] = i
				}
			}
			return stream
		}
		t.Run("BufferSize=10", func(t *testing.T) {
			runner := NewCVM(10, intTestComparator)
			for _, element := range newIncrementalDualTestStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 40_000)
		})

		t.Run("BufferSize=100", func(t *testing.T) {
			runner := NewCVM(100, intTestComparator)
			for _, element := range newIncrementalDualTestStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 20_000)
		})

		t.Run("BufferSize=1_000", func(t *testing.T) {
			runner := NewCVM(1_000, intTestComparator)
			for _, element := range newIncrementalDualTestStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 5_000)
		})

		t.Run("BufferSize=10_000", func(t *testing.T) {
			runner := NewCVM(10_000, intTestComparator)
			for _, element := range newIncrementalDualTestStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.InDelta(t, 50_000, runner.N(), 1_000)
		})

		t.Run("BufferSize=100_000", func(t *testing.T) {
			runner := NewCVM(100_000, intTestComparator)
			for _, element := range newIncrementalDualTestStream(1_000_000, 50_000) {
				runner.Process(element)
			}
			assert.Equal(t, 50_000, runner.N())
		})
	})

	// Disjointed blocks stream. Generated by at = xt + 10_000 * t / 10_000, , where xt is a random 4-digit number.
	// Thus it consists of ten thousand disjoint blocks of ten thousand numbers each. Expected number of distinct elements is 632_087 ≈ (1 − 1/e)m
	t.Run("DisjointBlocks", func(t *testing.T) {
		expected := 632_087
		newDisjointedBlocksTestStream := func() []int {
			total := 1_000_000
			stream := make([]int, total)
			for i := 0; i < total; i++ {
				x := rand.Intn(9_000) + 1_000
				stream[i] = x + (10_000 * i / 10_000)
			}
			return stream
		}

		t.Run("BufferSize=100", func(t *testing.T) {
			runner := NewCVM(100, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 200_000)
		})

		t.Run("BufferSize=1_000", func(t *testing.T) {
			runner := NewCVM(1_000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 100_000)
		})

		t.Run("BufferSize=10_000", func(t *testing.T) {
			runner := NewCVM(10_000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 20_000)
		})

		t.Run("BufferSize=100_000", func(t *testing.T) {
			runner := NewCVM(100_000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 5_000)
		})

		t.Run("BufferSize=1_000_000", func(t *testing.T) {
			runner := NewCVM(1000000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 1_500)
		})
	})

	// Disjointed blocks stream. Generated by at = xt + 10_000 * t / 10_000, , where xt is a random 4-digit number.
	// Thus it consists of ten thousand disjoint blocks of ten thousand numbers each. Expected number of distinct elements is 632_087 ≈ (1 − 1/e)m
	t.Run("IncrementalBinaryShifted", func(t *testing.T) {
		expected := 632_087
		newDisjointedBlocksTestStream := func() []int {
			total := 1_000_000
			stream := make([]int, total)
			for i := 0; i < total; i++ {
				x := rand.Intn(9_000) + 1_000
				stream[i] = x + (10_000 * i / 10_000)
			}
			return stream
		}

		t.Run("BufferSize=100", func(t *testing.T) {
			runner := NewCVM(100, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 200_000)
		})

		t.Run("BufferSize=1_000", func(t *testing.T) {
			runner := NewCVM(1_000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 100_000)
		})

		t.Run("BufferSize=10_000", func(t *testing.T) {
			runner := NewCVM(10_000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 20_000)
		})

		t.Run("BufferSize=100_000", func(t *testing.T) {
			runner := NewCVM(100_000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 5_000)
		})

		t.Run("BufferSize=1_000_000", func(t *testing.T) {
			runner := NewCVM(1000000, intTestComparator)
			for _, element := range newDisjointedBlocksTestStream() {
				runner.Process(element)
			}
			assert.InDelta(t, expected, runner.N(), 1_500)
		})
	})
}
