package main

import (
	"fmt"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	dq := deque.New[string]()
	dq.PushBack("apple", "banana", "cherry", "date")

	// Forward iteration
	fmt.Println("Forward iteration:")
	for v := range dq.Iterator() {
		fmt.Println(v)
	}

	// Reverse iteration
	fmt.Println("\nReverse iteration:")
	for v := range dq.DescendingeIterator() {
		fmt.Println(v)
	}
}
