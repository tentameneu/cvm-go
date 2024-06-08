package main

import "math/rand"

type cvm struct {
	stream []int
	buffer *treapBuffer
}

func (cvm *cvm) run() int {
	p := 1.0
	for _, a := range cvm.stream {
		cvm.buffer.delete(a)
		u := rand.Float64()
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
