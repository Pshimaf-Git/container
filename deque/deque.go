// package deque is a linked list based implementation of a double
// ended queue
package deque

import (
	"container/list"
	"errors"
	"fmt"
	"iter"
	"sync"
)

var (
	ErrTypeAssertion = errors.New("type assertion failed")
	ErrEmptyQueue    = errors.New("queue is empty")
)

// Deque represents a double-ended queue (deque) data structure
// that is thread-safe and generic over type T
type Deque[T any] struct {
	list *list.List
	mu   sync.RWMutex
}

// New creates and returns a new empty instance of Deque
func New[T any]() *Deque[T] {
	return &Deque[T]{
		list: list.New(),
	}
}

// zeroval returns the zero value for type T
func zeroval[T any]() T {
	var zero T
	return zero
}

// Len returns the number of elements in the deque
func (d *Deque[T]) Len() int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.list.Len()
}

// IsEmpty returns true if the deque contains no elements
func (d *Deque[T]) IsEmpty() bool {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.list.Len() == 0
}

// PushFront adds one or more values to the front of the deque
// in reverse order (last input becomes first in deque)
func (d *Deque[T]) PushFront(values ...T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i := len(values) - 1; i >= 0; i-- {
		d.list.PushFront(values[i])
	}
}

// PushBack appends one or more values to the end of the deque
// in the same order they were provided
func (d *Deque[T]) PushBack(values ...T) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, v := range values {
		d.list.PushBack(v)
	}
}

// PopFront removes and returns the first element from the deque.
// Returns an error if the deque is empty or type assertion fails.
func (d *Deque[T]) PopFront() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	const fancName = "(*Deque[T]).PopFront"

	if d.list.Len() == 0 {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrEmptyQueue)
	}

	elem := d.list.Front()
	defer d.list.Remove(elem)

	val, ok := elem.Value.(T)
	if !ok {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrTypeAssertion)
	}
	return val, nil
}

// PopBack removes and returns the last element from the deque.
// Returns an error if the deque is empty or type assertion fails.
func (d *Deque[T]) PopBack() (T, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	const fancName = "(*Deque[T]).PopBack"

	if d.list.Len() == 0 {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrEmptyQueue)
	}

	elem := d.list.Back()
	defer d.list.Remove(elem)

	val, ok := elem.Value.(T)
	if !ok {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrTypeAssertion)
	}
	return val, nil
}

// Front returns the first element from the deque without removing it.
// Returns an error if the deque is empty or type assertion fails.
func (d *Deque[T]) Front() (T, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	const fancName = "(*Deque[T]).Front"

	if d.list.Len() == 0 {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrEmptyQueue)
	}

	elem := d.list.Front()

	val, ok := elem.Value.(T)
	if !ok {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrTypeAssertion)
	}
	return val, nil
}

// Back returns the last element from the deque without removing it.
// Returns an error if the deque is empty or type assertion fails.
func (d *Deque[T]) Back() (T, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	const fancName = "(*Deque[T]).Back"

	if d.list.Len() == 0 {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrEmptyQueue)
	}

	elem := d.list.Back()

	val, ok := elem.Value.(T)
	if !ok {
		return zeroval[T](), fmt.Errorf("%s: %w", fancName, ErrTypeAssertion)
	}
	return val, nil
}

// Clear removes all elements from the deque and returns the count
// of elements that were removed
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

// ToArray converts the deque contents into a slice of type T
// Returns an empty slice if the deque is empty
func (d *Deque[T]) ToArray() []T {
	d.mu.RLock()
	defer d.mu.RUnlock()

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

// Get retrieves the element at the specified index without removing it.
// Returns the value and true if successful, zero value and false otherwise.
// The operation is optimized by traversing from the closer end (front or back).
func (d *Deque[T]) Get(index int) (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if index < 0 || index >= d.list.Len() || d.list.Len() == 0 {
		return zeroval[T](), false
	}

	var val T
	var ok bool

	if index < d.list.Len()/2 {
		e := d.list.Front()
		for i := 0; i < index; i++ {
			e = e.Next()
		}
		val, ok = e.Value.(T)

	} else {
		e := d.list.Back()
		for i := d.list.Len() - 1; i > index; i-- {
			e = e.Prev()
		}

		val, ok = e.Value.(T)
	}

	return val, ok
}

// Reverse reverses the order of elements in the deque in-place.
// If the deque is empty or has only one element, it does nothing
func (d *Deque[T]) Reverse() {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.list.Len() <= 1 {
		return
	}

	// In-place reversal without new list allocation
	front := d.list.Front()
	back := d.list.Back()
	for i := 0; i < d.list.Len()/2; i++ {
		front.Value, back.Value = back.Value, front.Value
		front = front.Next()
		back = back.Prev()
	}
}

// Count returns the number of occurrences of `target` in the deque.
// Uses the provided `equalFunc` to determine equality between elements
func (d *Deque[T]) Count(target T, equalFunc func(T, T) bool) int {
	d.mu.RLock()
	defer d.mu.RUnlock()

	count := 0
	for current := d.list.Front(); current != nil; current = current.Next() {
		if equalFunc(current.Value.(T), target) {
			count++
		}
	}
	return count
}

// Iterator returns a forward iterator (yields elements from front to back).
// The iterator terminates if the yield function returns false
func (d *Deque[T]) Iterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		d.mu.Lock()
		defer d.mu.Unlock()

		if d.list.Len() == 0 {
			return
		}

		index := 0
		for current := d.list.Front(); current != nil; current = current.Next() {
			if !yield(index, current.Value.(T)) {
				return
			}
			index++
		}
	}
}

// DescendingeIterator returns a reverse iterator (yields elements from back to front).
// The iterator terminates if the yield function returns false.
func (d *Deque[T]) DescendingeIterator() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		d.mu.Lock()
		defer d.mu.Unlock()

		if d.list.Len() == 0 {
			return
		}

		index := 0
		for current := d.list.Front(); current != nil; current = current.Next() {
			if !yield(index, current.Value.(T)) {
				return
			}
			index++
		}
	}
}

// Rotate rotates the deque by n positions.
// A positive n rotates elements to the right (toward the back),
// while a negative n rotates elements to the left (toward the front).
// If the deque is empty, has only one element, or n is 0, it does nothing.
// This operation is thread-safe
func (d *Deque[T]) Rotate(n int) {
	d.mu.Lock()
	defer d.mu.Unlock()

	length := d.list.Len()
	if d.list.Len() <= 1 || n == 0 {
		return
	}

	// Normalize n to be within [0, length)
	n = n % length
	if n < 0 {
		n += length
	}

	// Optimize by rotating in the most efficient direction
	if n <= length/2 {
		d.rotateRight(n)
	} else {
		d.rotateLeft(length - n)
	}

}

// rotateRight performs a right rotation by moving the last n elements
// to the front of the deque.
// Example: [1, 2, 3, 4] rotated right by 1 becomes [4, 1, 2, 3].
// Assumes n is positive and caller holds the lock
func (d *Deque[T]) rotateRight(n int) {
	for i := 0; i < n; i++ {
		d.list.MoveToFront(d.list.Back())
	}
}

// rotateLeft performs a left rotation by moving the first n elements
// to the back of the deque.
// Example: [1, 2, 3, 4] rotated left by 1 becomes [2, 3, 4, 1].
// Assumes n is positive and caller holds the lock
func (d *Deque[T]) rotateLeft(n int) {
	for i := n; i > 0; i-- {
		d.list.MoveToBack(d.list.Front())
	}
}
