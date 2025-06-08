// package deque is a linked list based implementation of a double
// ended queue
package deque

import (
	"container/list"
)

// structure of a dequeue
type Deque[T any] struct {
	list *list.List
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
	return d.list.Len()
}

// IsEmpty returns a boolean value that signals that the list is empty
func (d *Deque[T]) IsEmpty() bool {
	return d.Len() == 0
}

// PushFront adds the listed values ​​to the front of the queue
func (d *Deque[T]) PushFront(values ...T) {
	for i := len(values) - 1; i >= 0; i-- {
		d.list.PushFront(values[i])
	}
}

// PushBack adds the listed values ​​to the end of the queue
func (d *Deque[T]) PushBack(values ...T) {
	for _, v := range values {
		d.list.PushBack(v)
	}
}

// PopFront retrieves and removes the first element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
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
	if d.IsEmpty() {
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
	if d.IsEmpty() {
		return 0
	}

	cleared := 0

	for range d.Len() {
		d.list.Remove(d.list.Front())
		cleared++
	}

	return cleared
}

// ToArray converts all values ​​from a queue into an array
func (d *Deque[T]) ToArray() []T {
	if d.IsEmpty() {
		return []T{}
	}

	arr := make([]T, 0, d.Len())

	curr := d.list.Front()

	for curr.Next() != nil {
		val, ok := curr.Value.(T)
		if !ok {
			return arr
		}

		arr = append(arr, val)
		curr = curr.Next()
	}

	// Add last element value
	val, ok := curr.Value.(T)
	if !ok {
		return arr
	}

	arr = append(arr, val)

	return arr
}
