package stack

import "fmt"

type (
	Stack struct {
		top    *node
		length int
	}

	node struct {
		value    interface{}
		previous *node
	}
)

// New initialize a new stack
func New() *Stack {
	return &Stack{
		top:    nil,
		length: 0,
	}
}

// Push insert a new node in top of stack
func (slf *Stack) Push(value interface{}) {
	slf.top = &node{
		value:    value,
		previous: slf.top,
	}

	slf.length++
}

// Pop fetch top node from stack and then delete it
func (slf *Stack) Pop() interface{} {
	if slf.length == 0 {
		fmt.Println("stack is empty")
		return nil
	}

	topNode := slf.top
	slf.top = topNode.previous
	slf.length--

	return topNode.value
}

// Peek returns the top node's value and don't delete the top node
func (slf *Stack) Peek() interface{} {
	if slf.length == 0 {
		fmt.Println("stack is empty")
		return nil
	}
	return slf.top.value
}

// Size returns the size of stack
func (slf *Stack) Size() int {
	return slf.length
}
