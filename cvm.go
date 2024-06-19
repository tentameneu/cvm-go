package cvm

import (
	"math/rand"
)

// A CVM structure is used to run CVM algorithm to estimate number of distinct elements.
type CVM[T any] struct {
	buffer     *treapBuffer[T]
	bufferSize int
	total      int
	p          float64
}

// NewCVM returns new CVM struct with buffer of maximum size defined with bufferSize.
// Use comparator to define ordering of the elements.
func NewCVM[T any](bufferSize int, comparator Comparator[T]) *CVM[T] {
	return &CVM[T]{
		buffer:     newTreapBuffer(bufferSize, comparator),
		bufferSize: bufferSize,
		total:      0,
		p:          1.0,
	}
}

// N calculates estimated number of distinct elements using current buffer status.
func (cvm *CVM[T]) N() int {
	return int(float64(cvm.buffer.currentSize) / cvm.p)
}

// Process element from stream. Returns current estimated number of distinct elements using buffer status after processing element.
func (cvm *CVM[T]) Process(value T) int {
	cvm.total++
	u := rand.Float64()
	cvm.buffer.delete(value)

	if u >= cvm.p {
		return cvm.N()
	}
	if cvm.buffer.currentSize < cvm.buffer.maxSize {
		cvm.buffer.insert(newNode(value, u))
		return cvm.N()
	}
	if u > cvm.buffer.root.priority {
		cvm.p = u
	} else {
		cvm.p = cvm.buffer.root.priority
		cvm.buffer.delete(cvm.buffer.root.value)
		cvm.buffer.insert(newNode(value, u))
	}
	return cvm.N()
}
