package queue

type (
	Queue struct {
		head   *node
		length int
	}

	node struct {
		value interface{}
		next  *node
	}
)

// Size return the size of queue
func (slf *Queue) Size() int {
	return slf.length
}
