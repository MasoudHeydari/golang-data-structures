package main

import (
	"fmt"
	"queue/queue"
)

func main() {
	newQueue := queue.New()
	for i := 0; i < 10; i++ {
		newQueue.Enqueue(fmt.Sprintf("item: %v", i))
	}

	sizeOfQueue := newQueue.Size()
	for i := 0; i < sizeOfQueue; i++ {
		fmt.Println(newQueue.Dequeue())
	}
}
