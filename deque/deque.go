// package deque is a linked list based implementation of a double
// ended queue
package deque

import "container/list"

// PushBack — добавление в конец очереди.
// PushFront — добавление в начало очереди.
// PopBack — выборка из конца очереди.
// PopFront — выборка из начала очереди.
// IsEmpty — проверка наличия элементов.
// Clear — очистка.

// structure of a dequeue
type Deque struct {
	list *list.List
}

func New() *Deque {
	return &Deque{
		list: list.New(),
	}
}

func (d *Deque) Len() int {
	return d.list.Len()
}

func (d *Deque) IsEmpty() bool {
	return d.Len() == 0 || d == nil
}

func (d *Deque) PushFront(values ...any) {
	for _, v := range values {
		d.list.PushFront(v)
	}
}

func (d *Deque) PushBack(values ...any) {
	for _, v := range values {
		d.list.PushBack(v)
	}
}

func (d *Deque) PopFront() (any, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	return d.list.Front().Value, true
}

func (d *Deque) PopBack() (any, bool) {
	if d.IsEmpty() {
		return nil, false
	}

	return d.list.Back().Value, true
}

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

func (d *Deque) ToArray() []any {
	if d.IsEmpty() {
		return []any{}
	}

	arr := make([]any, 0, d.Len())

	for range d.Len() {
		arr = append(arr, d.list.Front().Value)
		d.list.Remove(d.list.Front())
	}

	return arr
}
