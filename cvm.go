package cvm

import (
	"math/rand"
)

type CVM struct {
	stream     []int
	buffer     *treapBuffer
	bufferSize int
}

func NewCVMRunner(stream []int, bufferSize int) *CVM {
	return &CVM{
		stream:     stream,
		buffer:     newTreapBuffer(bufferSize, func(x, y int) int { return x - y }),
		bufferSize: bufferSize,
	}
}

func (cvm *CVM) Run() int {
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
