package cvm

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tentameneu/cvm-go/internal/config"
	"github.com/tentameneu/cvm-go/internal/stream"
)

func newTestRepeatingStream(total, distinct int) []int {
	conf, _ := config.NewConfig(map[string]any{
		"genType":    "repeating",
		"total":      total,
		"distinct":   distinct,
		"bufferSize": total,
	})
	streamgenerator, _ := stream.NewStreamGenerator(conf)
	return streamgenerator.Generate()
}

func TestNew(t *testing.T) {
	t.Run("TreapBuffer", func(t *testing.T) {
		buffer := newTreapBuffer(10)
		assert.Nil(t, buffer.root)
		assert.Equal(t, 0, buffer.currentSize)
		assert.Equal(t, 10, buffer.maxSize)
	})

	t.Run("Node", func(t *testing.T) {
		node := newNode(123, 0.456)
		assert.Equal(t, 123, node.value)
		assert.Equal(t, 0.456, node.priority)
		assert.Nil(t, node.left)
		assert.Nil(t, node.right)
	})
}

func TestInsert(t *testing.T) {
	buffer := newTreapBuffer(10)

	t.Run("OnEmpty", func(t *testing.T) {
		buffer.insert(newNode(30, 0.200))
		assert.Equal(t, 1, buffer.currentSize)
		assert.Equal(t, 30, buffer.root.value)
		assert.Equal(t, 0.200, buffer.root.priority)
	})

	t.Run("NewRightLeaf", func(t *testing.T) {
		buffer.insert(newNode(40, 0.100))
		assert.Equal(t, 2, buffer.currentSize)
		assert.Equal(t, 30, buffer.root.value)
		assert.Equal(t, 0.200, buffer.root.priority)
		assert.Nil(t, buffer.root.left)
		assert.Equal(t, 40, buffer.root.right.value)
		assert.Equal(t, 0.100, buffer.root.right.priority)
	})

	t.Run("NewRoot", func(t *testing.T) {
		buffer.insert(newNode(20, 0.300))
		assert.Equal(t, 3, buffer.currentSize)
		assert.Equal(t, 20, buffer.root.value)
		assert.Equal(t, 0.300, buffer.root.priority)
		assert.Nil(t, buffer.root.left)
		assert.Equal(t, 30, buffer.root.right.value)
		assert.Equal(t, 0.200, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 40, buffer.root.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.priority)
	})

	t.Run("NewLeftLeaf", func(t *testing.T) {
		buffer.insert(newNode(10, 0.210))
		assert.Equal(t, 4, buffer.currentSize)
		assert.Equal(t, 20, buffer.root.value)
		assert.Equal(t, 0.300, buffer.root.priority)
		assert.Equal(t, 10, buffer.root.left.value)
		assert.Equal(t, 0.210, buffer.root.left.priority)
		assert.Equal(t, 30, buffer.root.right.value)
		assert.Equal(t, 0.200, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 40, buffer.root.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.priority)
	})

	t.Run("MiddleWithRotate", func(t *testing.T) {
		buffer.insert(newNode(15, 0.220))
		assert.Equal(t, 5, buffer.currentSize)
		assert.Equal(t, 20, buffer.root.value)
		assert.Equal(t, 0.300, buffer.root.priority)
		assert.Equal(t, 15, buffer.root.left.value)
		assert.Equal(t, 0.220, buffer.root.left.priority)
		assert.Equal(t, 10, buffer.root.left.left.value)
		assert.Equal(t, 0.210, buffer.root.left.left.priority)
		assert.Nil(t, buffer.root.left.right)
		assert.Equal(t, 30, buffer.root.right.value)
		assert.Equal(t, 0.200, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 40, buffer.root.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.priority)
	})

	t.Run("NewRootLowerValue", func(t *testing.T) {
		buffer.insert(newNode(18, 0.310))
		assert.Equal(t, 6, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 15, buffer.root.left.value)
		assert.Equal(t, 0.220, buffer.root.left.priority)
		assert.Equal(t, 20, buffer.root.right.value)
		assert.Equal(t, 0.300, buffer.root.right.priority)
		assert.Equal(t, 10, buffer.root.left.left.value)
		assert.Equal(t, 0.210, buffer.root.left.left.priority)
		assert.Nil(t, buffer.root.left.right)
		assert.Equal(t, 30, buffer.root.right.right.value)
		assert.Equal(t, 0.200, buffer.root.right.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 40, buffer.root.right.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.right.priority)
	})
}

func BenchmarkInsert(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000, 1_000_000}
	for _, length := range lengths {
		b.Run(fmt.Sprintf("%d", length), func(b *testing.B) {
			stream := newTestRepeatingStream(length, length)
			buffer := newTreapBuffer(length)
			b.ResetTimer()
			for _, element := range stream {
				buffer.insert(newNode(element, rand.Float64()))
			}
		})
	}
}

func TestInsertOverwrite(t *testing.T) {
	buffer := newTreapBuffer(10)
	buffer.insert(newNode(30, 0.200))
	buffer.insert(newNode(40, 0.100))
	buffer.insert(newNode(20, 0.300))
	buffer.insert(newNode(10, 0.210))
	buffer.insert(newNode(15, 0.220))
	buffer.insert(newNode(18, 0.310))

	newPriority1 := 0.230
	t.Run("InTheMiddle", func(t *testing.T) {
		buffer.insert(newNode(10, newPriority1))
		assert.Equal(t, 6, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 10, buffer.root.left.value)
		assert.Equal(t, newPriority1, buffer.root.left.priority)
		assert.Equal(t, 20, buffer.root.right.value)
		assert.Equal(t, 0.300, buffer.root.right.priority)
		assert.Nil(t, buffer.root.left.left)
		assert.Equal(t, 15, buffer.root.left.right.value)
		assert.Equal(t, 0.220, buffer.root.left.right.priority)
		assert.Equal(t, 30, buffer.root.right.right.value)
		assert.Equal(t, 0.200, buffer.root.right.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 40, buffer.root.right.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.right.priority)
	})

	newPriority2 := 0.350
	t.Run("UpgradeToRoot", func(t *testing.T) {
		buffer.insert(newNode(15, newPriority2))
		assert.Equal(t, 6, buffer.currentSize)
		assert.Equal(t, 15, buffer.root.value)
		assert.Equal(t, newPriority2, buffer.root.priority)
		assert.Equal(t, 10, buffer.root.left.value)
		assert.Equal(t, newPriority1, buffer.root.left.priority)
		assert.Equal(t, 18, buffer.root.right.value)
		assert.Equal(t, 0.310, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 20, buffer.root.right.right.value)
		assert.Equal(t, 0.300, buffer.root.right.right.priority)
		assert.Nil(t, buffer.root.right.right.left)
		assert.Equal(t, 30, buffer.root.right.right.right.value)
		assert.Equal(t, 0.200, buffer.root.right.right.right.priority)
		assert.Nil(t, buffer.root.right.right.right.left)
		assert.Equal(t, 40, buffer.root.right.right.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.right.right.priority)
	})

	newPriority3 := 0.150
	t.Run("DowngradeFromRoot", func(t *testing.T) {
		buffer.insert(newNode(15, newPriority3))
		assert.Equal(t, 6, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 10, buffer.root.left.value)
		assert.Equal(t, newPriority1, buffer.root.left.priority)
		assert.Nil(t, buffer.root.left.left)
		assert.Equal(t, 15, buffer.root.left.right.value)
		assert.Equal(t, newPriority3, buffer.root.left.right.priority)
		assert.Equal(t, 20, buffer.root.right.value)
		assert.Equal(t, 0.300, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Equal(t, 30, buffer.root.right.right.value)
		assert.Equal(t, 0.200, buffer.root.right.right.priority)
		assert.Nil(t, buffer.root.right.right.left)
		assert.Equal(t, 40, buffer.root.right.right.right.value)
		assert.Equal(t, 0.100, buffer.root.right.right.right.priority)
		assert.Nil(t, buffer.root.right.right.right.left)
	})
}

func BenchmarkInsertOverwrite(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000, 1_000_000}
	for _, length := range lengths {
		b.Run(fmt.Sprintf("%d", length), func(b *testing.B) {
			stream := newTestRepeatingStream(length, length/100)
			buffer := newTreapBuffer(length)
			for i := 0; i < length/100; i++ {
				buffer.insert(newNode(stream[i], rand.Float64()))
			}
			b.ResetTimer()
			for _, element := range stream {
				buffer.insert(newNode(element, rand.Float64()))
			}
		})
	}
}

func TestDelete(t *testing.T) {
	buffer := newTreapBuffer(10)
	buffer.insert(newNode(30, 0.200))
	buffer.insert(newNode(40, 0.100))
	buffer.insert(newNode(20, 0.300))
	buffer.insert(newNode(10, 0.210))
	buffer.insert(newNode(15, 0.220))
	buffer.insert(newNode(18, 0.310))

	t.Run("RightLeaf", func(t *testing.T) {
		buffer.delete(40)
		assert.Equal(t, 5, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 15, buffer.root.left.value)
		assert.Equal(t, 0.220, buffer.root.left.priority)
		assert.Equal(t, 20, buffer.root.right.value)
		assert.Equal(t, 0.300, buffer.root.right.priority)
		assert.Equal(t, 10, buffer.root.left.left.value)
		assert.Equal(t, 0.210, buffer.root.left.left.priority)
		assert.Nil(t, buffer.root.left.right)
		assert.Equal(t, 30, buffer.root.right.right.value)
		assert.Equal(t, 0.200, buffer.root.right.right.priority)
		assert.Nil(t, buffer.root.right.left)
		assert.Nil(t, buffer.root.right.right.right)
	})

	t.Run("LeftLeaf", func(t *testing.T) {
		buffer.delete(10)
		assert.Equal(t, 4, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 15, buffer.root.left.value)
		assert.Equal(t, 0.220, buffer.root.left.priority)
		assert.Equal(t, 20, buffer.root.right.value)
		assert.Equal(t, 0.300, buffer.root.right.priority)
		assert.Nil(t, buffer.root.left.left)
		assert.Nil(t, buffer.root.left.right)
		assert.Equal(t, 30, buffer.root.right.right.value)
		assert.Equal(t, 0.200, buffer.root.right.right.priority)
	})

	t.Run("Middle", func(t *testing.T) {
		buffer.delete(20)
		assert.Equal(t, 3, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 15, buffer.root.left.value)
		assert.Equal(t, 0.220, buffer.root.left.priority)
		assert.Equal(t, 30, buffer.root.right.value)
		assert.Equal(t, 0.200, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.right)
	})

	t.Run("NotExisting", func(t *testing.T) {
		buffer.delete(20)
		assert.Equal(t, 3, buffer.currentSize)
		assert.Equal(t, 18, buffer.root.value)
		assert.Equal(t, 0.310, buffer.root.priority)
		assert.Equal(t, 15, buffer.root.left.value)
		assert.Equal(t, 0.220, buffer.root.left.priority)
		assert.Equal(t, 30, buffer.root.right.value)
		assert.Equal(t, 0.200, buffer.root.right.priority)
		assert.Nil(t, buffer.root.right.right)
	})

	t.Run("Root", func(t *testing.T) {
		buffer.delete(18)
		assert.Equal(t, 2, buffer.currentSize)
		assert.Equal(t, 15, buffer.root.value)
		assert.Equal(t, 0.220, buffer.root.priority)
		assert.Nil(t, buffer.root.left)
		assert.Equal(t, 30, buffer.root.right.value)
		assert.Equal(t, 0.200, buffer.root.right.priority)
	})

	t.Run("LastLeafOnRoot", func(t *testing.T) {
		buffer.delete(30)
		assert.Equal(t, 1, buffer.currentSize)
		assert.Equal(t, 15, buffer.root.value)
		assert.Equal(t, 0.220, buffer.root.priority)
		assert.Nil(t, buffer.root.left)
		assert.Nil(t, buffer.root.right)
	})

	t.Run("LastRoot", func(t *testing.T) {
		buffer.delete(15)
		assert.Equal(t, 0, buffer.currentSize)
		assert.Nil(t, buffer.root)
	})
}

func BenchmarkContains(b *testing.B) {
	lengths := []int{1_000, 10_000, 100_000, 1_000_000}
	for _, length := range lengths {
		b.Run(fmt.Sprintf("%d", length), func(b *testing.B) {
			stream := newTestRepeatingStream(length, length-1)
			buffer := newTreapBuffer(length)
			for i := 0; i < length; i++ {
				buffer.insert(newNode(stream[i], rand.Float64()))
			}
			b.ResetTimer()
			for _, element := range stream {
				buffer.contains(element)
			}
		})
	}
}
