package main

import (
	"fmt"
	"sync"

	"github.com/Pshimaf-Git/container/deque"
)

func main() {
	dq := deque.New[int]()
	var wg sync.WaitGroup

	// Concurrent pushes
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			if val%2 == 0 {
				dq.PushFront(val)
			} else {
				dq.PushBack(val)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Deque length after concurrent pushes:", dq.Len())
	fmt.Println("Deque contents:", dq.ToArray())
}
