package main

import (
	"fmt"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	// Create a new deque
	dq := deque.New[int]()

	// Push elements to front and back
	dq.PushFront(1, 2, 3) // Deque: [3, 2, 1]
	dq.PushBack(4, 5)     // Deque: [3, 2, 1, 4, 5]

	// Check length
	fmt.Println("Length:", dq.Len()) // Output: 5

	// Pop elements
	front, _ := dq.PopFront()
	back, _ := dq.PopBack()
	fmt.Println("Front:", front) // Output: 3
	fmt.Println("Back:", back)   // Output: 5

	// Peek at elements
	currentFront, _ := dq.Front()
	currentBack, _ := dq.Back()
	fmt.Println("Current Front:", currentFront) // Output: 2
	fmt.Println("Current Back:", currentBack)   // Output: 4
}
