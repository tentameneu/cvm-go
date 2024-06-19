package cvm

import (
	"fmt"
	"io"
)

// Comparator is a function used to compare elements while saving them to a treap buffer.
// Function should return:
// return 0 if x == y.
// return < 0 (negative int) if x < y.
// return > 0 (positive int) if x > y.
type Comparator[T any] func(x, y T) int

type treapBuffer[T any] struct {
	root        *node[T]
	maxSize     int
	currentSize int
	comparator  Comparator[T]
}

type node[T any] struct {
	value    T
	priority float64
	left     *node[T]
	right    *node[T]
}

func newNode[T any](value T, priority float64) *node[T] {
	return &node[T]{
		value:    value,
		priority: priority,
		left:     nil,
		right:    nil,
	}
}

func newTreapBuffer[T any](maxSize int, comp Comparator[T]) *treapBuffer[T] {
	return &treapBuffer[T]{
		root:        nil,
		maxSize:     maxSize,
		currentSize: 0,
		comparator:  comp,
	}
}

func rightRotate[T any](node *node[T]) *node[T] {
	pivot := node.left
	temp := pivot.right
	pivot.right = node
	node.left = temp
	return pivot
}

func leftRotate[T any](node *node[T]) *node[T] {
	pivot := node.right
	temp := pivot.left
	pivot.left = node
	node.right = temp
	return pivot
}

func (tb *treapBuffer[T]) insert(newNode *node[T]) {
	if tb.contains((newNode.value)) {
		tb.delete(newNode.value)
	}

	tb.root = insertNode(tb.root, newNode, tb.comparator)
	tb.currentSize++
}

func insertNode[T any](root, newNode *node[T], comp Comparator[T]) *node[T] {
	if root == nil {
		return newNode
	}

	if comp(root.value, newNode.value) > 0 {
		root.left = insertNode(root.left, newNode, comp)
		if root.priority < newNode.priority {
			return rightRotate(root)
		}
		return root
	} else if comp(root.value, newNode.value) < 0 {
		root.right = insertNode(root.right, newNode, comp)
		if root.priority < newNode.priority {
			return leftRotate(root)
		}
		return root
	}

	return root
}

func (tb *treapBuffer[T]) delete(value T) {
	root, deleted := deleteNode(tb.root, value, tb.comparator, false)
	tb.root = root
	if deleted {
		tb.currentSize--
	}
}

func deleteNode[T any](root *node[T], value T, comp Comparator[T], found bool) (*node[T], bool) {
	if root == nil {
		return root, found
	}

	switch {
	case comp(value, root.value) < 0:
		root.left, found = deleteNode(root.left, value, comp, found)
	case comp(value, root.value) > 0:
		root.right, found = deleteNode(root.right, value, comp, found)
	case comp(value, root.value) == 0:
		switch {
		case root.left == nil:
			root = root.right
		case root.right == nil:
			root = root.left
		default:
			if root.left.priority < root.right.priority {
				root = leftRotate(root)
				root.left, _ = deleteNode(root.left, value, comp, found)
			} else {
				root = rightRotate(root)
				root.right, _ = deleteNode(root.right, value, comp, found)
			}
		}
		found = true
	}

	return root, found
}

func (tb *treapBuffer[T]) contains(value T) bool {
	current := tb.root

	for current != nil {
		if tb.comparator(value, current.value) == 0 {
			return true
		}

		if tb.comparator(value, current.value) < 0 {
			current = current.left
		} else {
			current = current.right
		}
	}

	return false
}

func (tb *treapBuffer[T]) printBasicInfo(writer io.Writer) {
	fmt.Fprintf(writer, "Size: %d\n", tb.currentSize)
	fmt.Fprint(writer, "Root: ")
	if tb.root != nil {
		printNode(writer, tb.root)
	} else {
		fmt.Fprintln(writer, tb.root)
	}
}

// func (tb *treapBuffer) printDetails() {
// 	fmt.Printf("Size: %d\n", tb.currentSize)
// 	fmt.Printf("Root: ")
// 	printNode(tb.root)
// 	printFrom((tb.root))
// }

// func printFrom(node *node) {
// 	if node != nil {
// 		printFrom(node.left)
// 		printNode(node)
// 		printFrom(node.right)
// 	}
// }

func printNode[T any](writer io.Writer, node *node[T]) {
	fmt.Fprintf(writer, "<Value: %v, Priority: %f>\n", node.value, node.priority)
}
