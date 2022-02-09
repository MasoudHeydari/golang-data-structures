package queue

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

// Enqueue adds a new node at the end of queue
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

// Size return the size of queue
func (slf *Queue) Size() int {
	return slf.length
}
