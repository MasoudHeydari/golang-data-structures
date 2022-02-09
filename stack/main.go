package main

import (
	"fmt"
	"stack/stack"
)

func main() {
	newStack := stack.New()
	for i := 0; i < 10; i++ {
		newStack.Push(fmt.Sprintf("item: %v", i))
	}

	sizeOfStack := newStack.Size()
	for i := 0; i < sizeOfStack; i++ {
		fmt.Println(newStack.Pop())
	}
}
