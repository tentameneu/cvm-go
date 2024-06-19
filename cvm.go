package cvm

import (
	"math/rand"
)

type CVM[T any] struct {
	stream     []T
	buffer     *treapBuffer[T]
	bufferSize int
}

func NewCVMRunner[T any](stream []T, bufferSize int, comparator Comparator[T]) *CVM[T] {
	return &CVM[T]{
		stream:     stream,
		buffer:     newTreapBuffer(bufferSize, comparator),
		bufferSize: bufferSize,
	}
}

func (cvm *CVM[T]) Run() int {
	p := 1.0
	for _, a := range cvm.stream {
		u := rand.Float64()
		cvm.buffer.delete(a)

		if u >= p {
			continue
		}
		if cvm.buffer.currentSize < cvm.buffer.maxSize {
			cvm.buffer.insert(newNode(a, u))
			continue
		}
		if u > cvm.buffer.root.priority {
			p = u
		} else {
			p = cvm.buffer.root.priority
			cvm.buffer.delete(cvm.buffer.root.value)
			cvm.buffer.insert(newNode(a, u))
		}
	}
	return int(float64(cvm.buffer.currentSize) / p)
}
