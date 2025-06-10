package main

import (
	"fmt"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	dq := deque.New[int]()
	dq.PushBack(1, 2, 3, 4, 5)

	fmt.Println("Original deque:", dq.ToArray()) // [1, 2, 3, 4, 5]

	// Rotate right by 2
	dq.Rotate(2)
	fmt.Println("After rotating right by 2:", dq.ToArray()) // [4, 5, 1, 2, 3]

	// Rotate left by 3
	dq.Rotate(-3)
	fmt.Println("After rotating left by 3:", dq.ToArray()) // [2, 3, 4, 5, 1]
}
