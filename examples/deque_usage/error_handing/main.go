package main

import (
	"fmt"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	dq := deque.New[string]()

	// Try to pop from empty deque
	val, err := dq.PopFront()
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: (*Deque[T]).PopFront: queue is empty
	} else {
		fmt.Println("Value:", val)
	}

	// Push some values
	dq.PushBack("hello", "world")

	// Try to get invalid index
	if val, ok := dq.Get(5); !ok {
		fmt.Println("Invalid index access") // Output: Invalid index access
	} else {
		fmt.Println("Value:", val)
	}
}
