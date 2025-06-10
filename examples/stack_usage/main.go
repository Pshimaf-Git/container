package main

import (
	"fmt"
	"os"

	"github.com/Pshimaf-Git/container/stack"
)

func main() {
	// Create new generic stack
	s := stack.New[int]()

	s.Push(1) // Pushed 1 into stack -> [1]
	s.Push(2) // -> [2,1]

	fmt.Println("Size after Push and before Pop:", s.Size())

	if v, ok := s.Pop(); !ok {
		fmt.Fprintln(os.Stderr, "ok is not true:", ok)
	} else {
		fmt.Println("Value:", v)
	}

	// After Pop -> [1]
	fmt.Println("Size after Pop:", s.Size())

	s.Pop()                    // Now stack is empty -> []
	if v, ok := s.Pop(); !ok { // ok is false
		fmt.Fprintln(os.Stderr, "ok is not true:", ok)
	} else {
		fmt.Println("Value:", v)
	}
}
