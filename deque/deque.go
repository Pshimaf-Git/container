// package deque is a linked list based implementation of a double
// ended queue
package deque

import (
	"container/list"
	"sync"
)

// structure of a dequeue
type Deque[T any] struct {
	list *list.List
	mu   sync.RWMutex
}

func New[T any]() *Deque[T] {
	return &Deque[T]{
		list: list.New(),
	}
}

func zeroval[T any]() T {
	var zero T
	return zero
}

// Len returns the total length of all elements in the queue
func (d *Deque[T]) Len() int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.list.Len()
}

// IsEmpty returns a boolean value that signals that the list is empty
func (d *Deque[T]) IsEmpty() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.list.Len() == 0
}

// PushFront adds the listed values ​​to the front of the queue
func (d *Deque[T]) PushFront(values ...T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i := len(values) - 1; i >= 0; i-- {
		d.list.PushFront(values[i])
	}
}

// PushBack adds the listed values ​​to the end of the queue
func (d *Deque[T]) PushBack(values ...T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, v := range values {
		d.list.PushBack(v)
	}
}

// PopFront retrieves and removes the first element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque[T]) PopFront() (T, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.list.Len() == 0 {
		return zeroval[T](), false
	}

	elem := d.list.Front()
	defer d.list.Remove(elem)

	val, ok := elem.Value.(T)
	if !ok {
		return zeroval[T](), false
	}

	return val, true
}

// PopBack retrieves and removes the last element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque[T]) PopBack() (T, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.list.Len() == 0 {
		return zeroval[T](), false
	}

	elem := d.list.Back()
	defer d.list.Remove(elem)

	val, ok := elem.Value.(T)
	if !ok {
		return zeroval[T](), false
	}

	return val, true
}

// Clear removes all elements from queue
func (d *Deque[T]) Clear() int {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.list.Len() == 0 {
		return 0
	}

	cleared := 0

	for e := d.list.Front(); e != nil; {
		next := e.Next()
		d.list.Remove(e)
		e = next
		cleared++
	}

	return cleared
}

// ToArray converts all values ​​from a queue into an array
func (d *Deque[T]) ToArray() []T {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.list.Len() == 0 {
		return []T{}
	}

	arr := make([]T, 0, d.list.Len())

	for e := d.list.Front(); e != nil; e = e.Next() {
		val, ok := e.Value.(T)
		if !ok {
			break
		}
		arr = append(arr, val)
	}

	return arr
}
