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

// New create a new queue
func New() *Queue {
	return &Queue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

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

// Dequeue fetch the head node from the queue and delete it
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

// Peek fetch the head node from the queue without removing it
func (slf *Queue) Peek() interface{} {
	if slf.length == 0 {
		fmt.Println("queue is empty")
		return nil
	}

	return slf.head.value
}

// Size return the size of queue
func (slf *Queue) Size() int {
	return slf.length
}
