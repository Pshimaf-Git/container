package main

import (
	"fmt"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	dq := deque.New[int]()
	dq.PushBack(1, 2, 3, 2, 4, 2, 5)

	// Count occurrences of 2
	count := dq.Count(2, func(a, b int) bool { return a == b })
	fmt.Println("Number of 2s:", count) // Output: 3
}
