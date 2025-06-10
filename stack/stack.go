// Package stack provides a lock-free, thread-safe generic stack implementation.
//
// The stack supports concurrent push, pop and size operations using atomic operations
// rather than mutex locks, which can provide better performance in high-concurrency
// scenarios. The implementation is based on a linked list using unsafe.Pointers
// and atomic compare-and-swap (CAS) operations.
//
// The stack follows LIFO (Last-In-First-Out) semantics and is generic, working
// with any type. Empty stack pops return the type's zero value.
//
// Example usage:
//
//	s := stack.New[int]()
//	s.Push(42)
//	val, ok := s.Pop()  // returns 42, true
//	val, ok = s.Pop()   // returns 0, false (empty stack)
package stack

import (
	"sync/atomic"
	"unsafe"
)

// item represents a single element in the stack, containing a value and a pointer to the
// next item
type item[T any] struct {
	value T
	next  unsafe.Pointer
}

// Stack is a thread-safe, generic LIFO (Last-In-First-Out) data structure implemented
// using atomic operations.
// It supports concurrent push and pop operations without locks
type Stack[T any] struct {
	head unsafe.Pointer
	size atomic.Uint32
}

// New creates and returns a new, empty Stack for type T.
// The stack is initialized with a nil head pointer
func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() uint32 {
	return s.size.Load()
}

func (s *Stack[T]) Empty() bool {
	return s.size.Load() == 0
}

// Push adds a new value to the top of the stack.
// This operation is atomic and thread-safe, using compare-and-swap (CAS) to handle
// concurrent access.
// The value will be the first one to be returned by subsequent Pop operations
func (s *Stack[T]) Push(value T) {
	node := &item[T]{value: value}

	for {
		head := atomic.LoadPointer(&s.head)
		node.next = head

		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(node)) {
			s.size.Add(1)
			return
		}
	}
}

// Pop removes and returns the value at the top of the stack.
// It returns the value and a boolean indicating whether the operation was successful.
// If the stack is empty, it returns the zero value of type T and false.
// This operation is atomic and thread-safe, using compare-and-swap (CAS) to handle
// concurrent access
func (s *Stack[T]) Pop() (T, bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return zeroval[T](), false
		}

		next := atomic.LoadPointer(&(*item[T])(head).next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {

			// Decrement length
			s.size.Add(^uint32(0))

			return (*item[T])(head).value, true
		}
	}
}

// zeroval returns the zero value for type T.
// This is used to return a valid value when popping from an empty stack
func zeroval[T any]() T {
	var z T
	return z
}
