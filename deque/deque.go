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

// New returns new instance of Deque struct
func New[T any]() *Deque[T] {
	return &Deque[T]{
		list: list.New(),
	}
}

// zeroval return zero value of type T (helper func)
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
	return val, ok
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
	return val, ok
}

// Front retrieves the first element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque[T]) Front() (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.list.Len() == 0 {
		return zeroval[T](), false
	}

	e := d.list.Front()
	if e == nil {
		return zeroval[T](), false
	}

	val, ok := e.Value.(T)
	return val, ok
}

// Back retrieves the last element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque[T]) Back() (T, bool) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if d.list.Len() == 0 {
		return zeroval[T](), false
	}

	e := d.list.Back()
	if e == nil {
		return zeroval[T](), false
	}

	val, ok := e.Value.(T)
	return val, ok
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

// Get returns the value of the element from the list under the passed index, the boolean
// value signals the success of the retrieval
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

// Remove removes e from d.list if e is an element of list d.list. It returns the element
// value e.Value. The element must not be nil and must belong to this deque.
func (d *Deque[T]) Remove(e *list.Element) (T, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if e == nil || d.list.Len() == 0 {
		return zeroval[T](), false
	}

	// Optimization: choosing which end to start the search from
	pos := getElementPosition(d.list, e)
	if pos < 0 {
		return zeroval[T](), false
	}

	found := false

	if pos < d.list.Len()/2 {
		// Find from start
		for current := d.list.Front(); current != nil; current = current.Next() {
			if current == e {
				found = true
				break
			}
		}
	} else {
		// Find from end
		for current := d.list.Back(); current != nil; current = current.Prev() {
			if current == e {
				found = true
				break
			}
		}
	}

	if found {
		val, ok := e.Value.(T)
		d.list.Remove(e)
		return val, ok
	}

	return zeroval[T](), false
}

// Helper function for determining the position of an element
func getElementPosition(l *list.List, e *list.Element) int {
	pos := 0
	for current := l.Front(); current != nil; current = current.Next() {
		if current == e {
			return pos
		}
		pos++
	}
	return -1
}
