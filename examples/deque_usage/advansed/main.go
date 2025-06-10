package main

import (
	"fmt"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	dq := deque.New[float64]()

	// Fill deque
	dq.PushBack(1.1, 2.2, 3.3, 4.4, 5.5)

	// Reverse the deque
	dq.Reverse()
	fmt.Println("Reversed:", dq.ToArray()) // [5.5, 4.4, 3.3, 2.2, 1.1]

	// Clear the deque
	cleared := dq.Clear()
	fmt.Println("Cleared", cleared, "elements") // Output: Cleared 5 elements
	fmt.Println("Is empty now:", dq.IsEmpty())  // Output: true
}
