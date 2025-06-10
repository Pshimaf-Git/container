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
	for i, v := range dq.Iterator() {
		fmt.Printf("Index: %d Value: %v\n", i, v)
	}

	// Reverse iteration
	fmt.Println("\nReverse iteration:")
	for i, v := range dq.DescendingeIterator() {
		fmt.Printf("Index: %d Value: %v\n", i, v)
	}
}
