// package deque is a linked list based implementation of a double
// ended queue
package deque

import (
	"container/list"
)

// structure of a dequeue
type Deque struct {
	list *list.List
}

func New() *Deque {
	return &Deque{
		list: list.New(),
	}
}

// Len returns the total length of all elements in the queue
func (d *Deque) Len() int {
	return d.list.Len()
}

// IsEmpty returns a boolean value that signals that the list is empty
func (d *Deque) IsEmpty() bool {
	return d.Len() == 0 || d == nil
}

// PushFront adds the listed values ​​to the front of the queue
func (d *Deque) PushFront(values ...any) {
	values = revers(values)
	for _, v := range values {
		d.list.PushFront(v)
	}
}

// PushBack adds the listed values ​​to the end of the queue
func (d *Deque) PushBack(values ...any) {
	for _, v := range values {
		d.list.PushBack(v)
	}
}

// PopFront retrieves and removes the first element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque) PopFront() (any, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	elem := d.list.Front()
	defer d.list.Remove(elem)

	return elem.Value, true
}

// PopBack retrieves and removes the last element from the queue, a
// boolean value signals that the resulting value is not a zero value
func (d *Deque) PopBack() (any, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	elem := d.list.Back()
	defer d.list.Remove(elem)

	return elem.Value, true
}

// Clear removes all elements from queue
func (d *Deque) Clear() int {
	cleared := 0
	if d.IsEmpty() {
		return cleared
	}

	for range d.Len() {
		d.list.Remove(d.list.Front())
		cleared++
	}

	return cleared
}

// ToArray converts all values ​​from a queue into an array
func (d *Deque) ToArray() []any {
	if d.IsEmpty() {
		return []any{}
	}

	arr := make([]any, 0, d.Len())

	curr := d.list.Front()

	for curr.Next() != nil {
		arr = append(arr, curr.Value)
		curr = curr.Next()
	}

	// Add last element value
	arr = append(arr, curr.Value)

	return arr
}

// reverse returns a reversed copy of the input slide
func revers(a []any) []any {
	if len(a) == 0 {
		return a
	}

	res := a
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
