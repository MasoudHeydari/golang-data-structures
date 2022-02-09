package queue

import "fmt"

type (
	Queue struct {
		head, tail *node
		length     int
	}

	node struct {
		value interface{}
		next  *node
	}
)

// Enqueue insert a new node at the end of queue
func (slf *Queue) Enqueue(newValue interface{}) {
	newNode := &node{
		value: newValue,
		next:  nil,
	}
	if slf.length == 0 {
		// queue is empty
		slf.head = newNode
		slf.tail = newNode
	} else {
		// queue has at least one node
		slf.tail.next = newNode
		slf.tail = newNode
	}

	slf.length++
}

// Dequeue fetch a node from the head of queue and delete it from queue
func (slf *Queue) Dequeue() interface{} {
	if slf.length == 0 {
		fmt.Println("queue is empty")
		return nil
	}

	headNode := slf.head
	slf.head = headNode.next
	slf.length--

	if slf.length == 1 {
		slf.tail = nil
	}

	return headNode.value
}

// Size return the size of queue
func (slf *Queue) Size() int {
	return slf.length
}
