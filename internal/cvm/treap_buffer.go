package cvm

import "fmt"

type treapBuffer struct {
	root        *node
	maxSize     int
	currentSize int
}

type node struct {
	value    int
	priority float64
	left     *node
	right    *node
}

func newNode(value int, priority float64) *node {
	return &node{
		value:    value,
		priority: priority,
		left:     nil,
		right:    nil,
	}
}

func newTreapBuffer(maxSize int) *treapBuffer {
	return &treapBuffer{
		root:        nil,
		maxSize:     maxSize,
		currentSize: 0,
	}
}

func rightRotate(node *node) *node {
	pivot := node.left
	temp := pivot.right
	pivot.right = node
	node.left = temp
	return pivot
}

func leftRotate(node *node) *node {
	pivot := node.right
	temp := pivot.left
	pivot.left = node
	node.right = temp
	return pivot
}

func (tb *treapBuffer) insert(newNode *node) {
	if tb.contains((newNode.value)) {
		tb.delete(newNode.value)
	}

	tb.root = insertNode(tb.root, newNode)
	tb.currentSize++
}

func insertNode(root, newNode *node) *node {
	if root == nil {
		return newNode
	}

	if root.value > newNode.value {
		root.left = insertNode(root.left, newNode)
		if root.priority < newNode.priority {
			return rightRotate(root)
		}
		return root
	} else if root.value < newNode.value {
		root.right = insertNode(root.right, newNode)
		if root.priority < newNode.priority {
			return leftRotate(root)
		}
		return root
	}

	return root
}

func (tb *treapBuffer) delete(value int) {
	root, deleted := deleteNode(tb.root, value, false)
	tb.root = root
	if deleted {
		tb.currentSize--
	}
}

func deleteNode(root *node, value int, found bool) (*node, bool) {
	if root == nil {
		return root, found
	}

	switch {
	case value < root.value:
		root.left, found = deleteNode(root.left, value, found)
	case value > root.value:
		root.right, found = deleteNode(root.right, value, found)
	case value == root.value:
		switch {
		case root.left == nil:
			root = root.right
		case root.right == nil:
			root = root.left
		default:
			if root.left.priority < root.right.priority {
				root = leftRotate(root)
				root.left, _ = deleteNode(root.left, value, found)
			} else {
				root = rightRotate(root)
				root.right, _ = deleteNode(root.right, value, found)
			}
		}
		found = true
	}

	return root, found
}

func (tb *treapBuffer) contains(value int) bool {
	current := tb.root

	for current != nil {
		if value == current.value {
			return true
		}

		if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
	}

	return false
}

func (tb *treapBuffer) printBasicInfo() {
	fmt.Printf("Size: %d\n", tb.currentSize)
	fmt.Printf("Root: ")
	printNode(tb.root)
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

func printNode(node *node) {
	fmt.Printf("<Value: %d, Priority: %f>\n", node.value, node.priority)
}
