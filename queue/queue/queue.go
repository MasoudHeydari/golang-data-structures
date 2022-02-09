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
