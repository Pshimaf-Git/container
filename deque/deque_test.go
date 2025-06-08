package deque

import (
	"container/list"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	d := New[int]()
	assert.NotNil(t, d)
	assert.NotNil(t, d.list)
	assert.Equal(t, 0, d.Len())
	assert.True(t, d.IsEmpty())
}

func TestPushFront(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty input", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[int]()
			d.PushFront(tt.input...)
			assert.Equal(t, len(tt.expected), d.Len())

			// Verify order
			arr := d.ToArray()
			assert.Equal(t, tt.expected, arr)
		})
	}
}

func TestPushBack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"single element", []int{1}, []int{1}},
		{"multiple elements", []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty input", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[int]()
			d.PushBack(tt.input...)
			assert.Equal(t, len(tt.expected), d.Len())

			// Verify order
			arr := d.ToArray()
			assert.Equal(t, tt.expected, arr)
		})
	}
}

func TestPopFront(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *Deque[int]
		expectedVal int
		expectedErr error
	}{
		{
			"empty deque",
			func() *Deque[int] { return New[int]() },
			0,
			ErrEmprtQueue,
		},
		{
			"single element",
			func() *Deque[int] {
				d := New[int]()
				d.PushBack(42)
				return d
			},
			42,
			nil,
		},
		{
			"multiple elements",
			func() *Deque[int] {
				d := New[int]()
				d.PushBack(1)
				d.PushBack(2)
				d.PushBack(3)
				return d
			},
			1,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup()
			val, err := d.PopFront()

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedVal, val)
		})
	}
}

func TestPopBack(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *Deque[int]
		expectedVal int
		expectedErr error
	}{
		{
			"empty deque",
			func() *Deque[int] { return New[int]() },
			0,
			ErrEmprtQueue,
		},
		{
			"single element",
			func() *Deque[int] {
				d := New[int]()
				d.PushBack(42)
				return d
			},
			42,
			nil,
		},
		{
			"multiple elements",
			func() *Deque[int] {
				d := New[int]()
				d.PushBack(1)
				d.PushBack(2)
				d.PushBack(3)
				return d
			},
			3,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup()
			val, err := d.PopBack()

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedVal, val)
		})
	}
}

func TestFront(t *testing.T) {
	d := New[int]()
	_, err := d.Front()
	assert.ErrorIs(t, err, ErrEmprtQueue)

	d.PushBack(1)
	d.PushBack(2)
	val, err := d.Front()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	assert.Equal(t, 2, d.Len()) // Shouldn't remove the element
}

func TestBack(t *testing.T) {
	d := New[int]()
	_, err := d.Back()
	assert.ErrorIs(t, err, ErrEmprtQueue)

	d.PushBack(1)
	d.PushBack(2)
	val, err := d.Back()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	assert.Equal(t, 2, d.Len()) // Shouldn't remove the element
}

func TestClear(t *testing.T) {
	d := New[int]()
	assert.Equal(t, 0, d.Clear())

	d.PushBack(1)
	d.PushBack(2)
	assert.Equal(t, 2, d.Clear())
	assert.True(t, d.IsEmpty())
}

func TestToArray(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *Deque[int]
		expected []int
	}{
		{
			"empty deque",
			func() *Deque[int] { return New[int]() },
			[]int{},
		},
		{
			"single element",
			func() *Deque[int] {
				d := New[int]()
				d.PushBack(1)
				return d
			},
			[]int{1},
		},
		{
			"multiple elements",
			func() *Deque[int] {
				d := New[int]()
				d.PushBack(1)
				d.PushBack(2)
				d.PushBack(3)
				return d
			},
			[]int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.setup()
			assert.Equal(t, tt.expected, d.ToArray())
		})
	}
}

func TestGet(t *testing.T) {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	tests := []struct {
		name     string
		index    int
		expected int
		ok       bool
	}{
		{"first element", 0, 1, true},
		{"middle element", 1, 2, true},
		{"last element", 2, 3, true},
		{"negative index", -1, 0, false},
		{"index out of range", 3, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, ok := d.Get(tt.index)
			assert.Equal(t, tt.expected, val)
			assert.Equal(t, tt.ok, ok)
		})
	}
}

func TestRemove(t *testing.T) {
	d := New[int]()
	d.PushBack(1)
	d.PushBack(2)
	d.PushBack(3)

	// Get the middle element
	e := d.list.Front().Next()

	// Test successful removal
	val, err := d.Remove(e)
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	assert.Equal(t, 2, d.Len())

	// Try to remove nil element
	_, err = d.Remove(nil)
	assert.ErrorIs(t, err, ErrNilElem)
	assert.Contains(t, err.Error(), "(*Deque[T]).Remove")

	// Try to remove element not in the list
	newList := New[int]()
	newList.PushBack(99)
	_, err = d.Remove(newList.list.Front())
	assert.ErrorIs(t, err, ErrNotFoundElem)
	assert.Contains(t, err.Error(), "(*Deque[T]).Remove")

	// Try to remove from empty deque
	d.Clear()
	_, err = d.Remove(e) // e is no longer valid
	assert.ErrorIs(t, err, ErrEmprtQueue)
	assert.Contains(t, err.Error(), "(*Deque[T]).Remove")
}

func TestConcurrency(t *testing.T) {
	d := New[int]()
	const numWorkers = 100
	const numOperations = 1000

	// Worker function that performs various operations
	worker := func(id int, done chan struct{}) {
		for i := 0; i < numOperations; i++ {
			// Alternate between different operations
			switch i % 4 {
			case 0:
				d.PushFront(id*1000 + i)
			case 1:
				d.PushBack(id*1000 + i)
			case 2:
				d.PopFront()
			case 3:
				d.PopBack()
			}
		}
		done <- struct{}{}
	}

	// Start workers
	done := make(chan struct{})
	for i := 0; i < numWorkers; i++ {
		go worker(i, done)
	}

	// Wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		<-done
	}

	// Verify the deque is in a consistent state
	// We can't predict the exact contents due to concurrency,
	// but we can check basic invariants
	assert.GreaterOrEqual(t, d.Len(), 0)

	// Check that all elements can be accessed without panic
	for i := 0; i < d.Len(); i++ {
		val, ok := d.Get(i)
		if ok {
			assert.NotNil(t, val)
		}
	}
}

func TestDeque_ContainsElement(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() (*Deque[int], *list.Element, bool)
		expected bool
	}{
		{
			name: "element exists",
			setup: func() (*Deque[int], *list.Element, bool) {
				d := New[int]()
				e := d.list.PushBack(1)
				d.list.PushBack(2)
				return d, e, true
			},
			expected: true,
		},
		{
			name: "element doesn't exist",
			setup: func() (*Deque[int], *list.Element, bool) {
				d := New[int]()
				d.list.PushBack(1)
				d.list.PushBack(2)
				return d, &list.Element{Value: 3}, false
			},
			expected: false,
		},
		{
			name: "empty deque",
			setup: func() (*Deque[int], *list.Element, bool) {
				return New[int](), &list.Element{Value: 1}, false
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, e, _ := tt.setup()
			assert.Equal(t, tt.expected, d.ContainsElement(e))
		})
	}
}

func TestDeque_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "empty deque",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "two elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[int]()
			for _, v := range tt.input {
				d.list.PushBack(v)
			}

			d.Reverse()
			assert.Equal(t, tt.expected, d.ToArray())
		})
	}
}

func TestDeque_Iterator(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "normal iteration",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty deque",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[int]()
			for _, v := range tt.input {
				d.list.PushBack(v)
			}

			var result = []int{}
			for v := range d.Iterator() {
				result = append(result, v)
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDeque_DescendingIterator(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "normal reverse iteration",
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			name:     "empty deque",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[int]()
			for _, v := range tt.input {
				d.list.PushBack(v)
			}

			var result = []int{}
			for v := range d.DescendingeIterator() {
				result = append(result, v)
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDeque_Rotate(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		n        int
		expected []int
	}{
		{
			name:     "rotate right by 1",
			input:    []int{1, 2, 3, 4},
			n:        1,
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "rotate left by 1",
			input:    []int{1, 2, 3, 4},
			n:        -1,
			expected: []int{2, 3, 4, 1},
		},
		{
			name:     "rotate by length (no change)",
			input:    []int{1, 2, 3, 4},
			n:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "rotate by multiple lengths",
			input:    []int{1, 2, 3, 4},
			n:        5, // same as rotate by 1
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "empty deque",
			input:    []int{},
			n:        1,
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			n:        1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[int]()
			for _, v := range tt.input {
				d.list.PushBack(v)
			}

			d.Rotate(tt.n)
			assert.Equal(t, tt.expected, d.ToArray())
		})
	}
}
func TestConcurrentAccess(t *testing.T) {
	d := New[int]()
	stop := make(chan struct{})

	// Writer goroutine
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				d.PushBack(1)
				d.PushFront(2)
				time.Sleep(time.Millisecond)
			}
		}
	}()

	// Reader goroutine
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				d.Front()
				d.Back()
				d.Len()
				d.IsEmpty()
				time.Sleep(time.Millisecond)
			}
		}
	}()

	// Remover goroutine
	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				d.PopFront()
				d.PopBack()
				time.Sleep(time.Millisecond)
			}
		}
	}()

	// Let them run for a while
	time.Sleep(100 * time.Millisecond)
	close(stop)

	// Final check - deque should be in consistent state
	assert.GreaterOrEqual(t, d.Len(), 0)
}
