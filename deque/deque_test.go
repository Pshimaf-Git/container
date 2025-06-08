// package deque is a linked list based implementation of a double

// ended queue

package deque

import (
	"container/list"
	"reflect"
	"testing"
)

func TestDeque_Len(t *testing.T) {
	tests := []struct {
		name       string
		initValues []any
		want       int
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       0,
		},

		{
			name:       "full deque",
			initValues: []any{1, 2, "Hello", false},
			want:       4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			if got := d.Len(); got != tt.want {
				t.Errorf("Deque.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_IsEmpty(t *testing.T) {
	tests := []struct {
		name       string
		initValues []any
		want       bool
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       true,
		},

		{
			name:       "full deque",
			initValues: []any{1, 2, 3, 4, 5},
			want:       false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			if got := d.IsEmpty(); got != tt.want {
				t.Errorf("Deque.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_ToArray(t *testing.T) {
	tests := []struct {
		name       string
		initValues []any
		want       []any
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       []any{},
		},

		{
			name:       "mixed types",
			initValues: []any{1, "hello", 0x10},
			want:       []any{1, "hello", 0x10},
		},

		{
			name:       "base-case",
			initValues: []any{1, '1'},
			want:       []any{1, '1'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			if got := d.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.ToArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PushFront(t *testing.T) {
	type args struct {
		values []any
	}

	tests := []struct {
		name       string
		args       args
		initValues []any
		want       []any
	}{
		{
			name:       "push one to empty deque",
			initValues: []any{},
			args: args{
				values: []any{1},
			},
			want: []any{1},
		},

		{
			name:       "push one to full deque",
			initValues: []any{1, true, "hello"},
			args: args{
				values: []any{0},
			},
			want: []any{0, 1, true, "hello"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			d.PushFront(tt.args.values...)

			if got := d.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.PushFront(%v) got = %v, want = %v", tt.args.values, got, tt.want)
			}
		})
	}
}

func TestDeque_PushBack(t *testing.T) {
	type args struct {
		values []any
	}

	tests := []struct {
		name       string
		args       args
		initValues []any
		want       []any
	}{
		{
			name:       "push one to empty deque",
			initValues: []any{},
			args: args{
				values: []any{1},
			},
			want: []any{1},
		},

		{
			name:       "push one to full deque",
			initValues: []any{1, true, "hello"},
			args: args{
				values: []any{0},
			},
			want: []any{1, true, "hello", 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			d.PushBack(tt.args.values...)

			if got := d.ToArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.PushBack(%v) got = %v, want = %v", tt.args.values, got, tt.want)
			}
		})
	}

}

func TestDeque_Clear(t *testing.T) {
	tests := []struct {
		name       string
		initValues []any
		want       int
	}{
		{
			name:       "empty deque1",
			initValues: []any{},
			want:       0,
		},

		{
			name:       "full deque",
			initValues: []any{1, 2, 3},
			want:       3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			if got := d.Clear(); got != tt.want || d.list.Len() != 0 {
				t.Errorf("Deque.Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PopFront(t *testing.T) {
	tests := []struct {
		name       string
		initValues []any
		want       any
		want1      bool
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       nil,
			want1:      false,
		},

		{
			name:       "full deque",
			initValues: []any{1, 2, 3, "hello", &list.Element{}},
			want:       1,
			want1:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			got, got1 := d.PopFront()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.PopFront() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.PopFront() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeque_PopBack(t *testing.T) {
	tests := []struct {
		name       string
		initValues []any
		want       any
		want1      bool
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       nil,
			want1:      false,
		},

		{
			name:       "full deque",
			initValues: []any{1, true, &testing.B{}, ' '},
			want:       ' ',
			want1:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			got, got1 := d.PopBack()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.PopBack() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.PopBack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeque_Front(t *testing.T) {
	testCases := []struct {
		name       string
		initValues []any
		want       any
		want1      bool
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       nil,
			want1:      false,
		},

		{
			name:       "full deqque",
			initValues: []any{1, 32.2},
			want:       1,
			want1:      true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			got, got1 := d.Front()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.Front() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.Front() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeque_Back(t *testing.T) {
	testCases := []struct {
		name       string
		initValues []any
		want       any
		want1      bool
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			want:       nil,
			want1:      false,
		},

		{
			name:       "full deqque",
			initValues: []any{1, true},
			want:       true,
			want1:      true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			got, got1 := d.Back()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.Back() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.Back() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDeque_Get(t *testing.T) {
	type args struct {
		index int
	}

	testCases := []struct {
		name       string
		initValues []any
		args       args
		want       any
		want1      bool
	}{
		{
			name:       "empty deque",
			initValues: []any{},
			args: args{
				index: 0,
			},
			want:  nil,
			want1: false,
		},

		{
			name:       "negative index",
			initValues: []any{1, 2, "hello"},
			args: args{
				index: -1,
			},
			want:  nil,
			want1: false,
		},

		{
			name:       "index out of range",
			initValues: []any{1, 2, 3},
			args: args{
				index: 1000,
			},
			want:  nil,
			want1: false,
		},

		{
			name:       "valid index",
			initValues: []any{1, 2},
			args: args{
				index: 0,
			},
			want:  1,
			want1: true,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			d := New[any]()
			d.PushFront(tt.initValues...)

			got, got1 := d.Get(tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.Get(%v) got = %v, want %v", tt.args.index, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.Get(%v) got1 = %v, want %v", tt.args.index, got1, tt.want1)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name        string
		setup       func() (*Deque[int], *list.Element)
		expectValue int
		expectOk    bool
		expectLen   int
	}{
		{
			name: "Remove from empty deque",
			setup: func() (*Deque[int], *list.Element) {
				d := New[int]()
				return d, &list.Element{Value: 42}
			},
			expectValue: 0,
			expectOk:    false,
			expectLen:   0,
		},
		{
			name: "Remove nil element",
			setup: func() (*Deque[int], *list.Element) {
				d := New[int]()
				d.PushBack(1, 2, 3)
				return d, nil
			},
			expectValue: 0,
			expectOk:    false,
			expectLen:   3,
		},
		{
			name: "Remove front element",
			setup: func() (*Deque[int], *list.Element) {
				d := New[int]()
				d.PushBack(1, 2, 3)
				return d, d.list.Front()
			},
			expectValue: 1,
			expectOk:    true,
			expectLen:   2,
		},
		{
			name: "Remove middle element",
			setup: func() (*Deque[int], *list.Element) {
				d := New[int]()
				d.PushBack(1, 2, 3)
				return d, d.list.Front().Next()
			},
			expectValue: 2,
			expectOk:    true,
			expectLen:   2,
		},
		{
			name: "Remove back element",
			setup: func() (*Deque[int], *list.Element) {
				d := New[int]()
				d.PushBack(1, 2, 3)
				return d, d.list.Back()
			},
			expectValue: 3,
			expectOk:    true,
			expectLen:   2,
		},
		{
			name: "Remove non-existent element from another deque",
			setup: func() (*Deque[int], *list.Element) {
				d1 := New[int]()
				d1.PushBack(1, 2, 3)

				d2 := New[int]()
				d2.PushBack(4, 5, 6)
				return d1, d2.list.Front()
			},
			expectValue: 0,
			expectOk:    false,
			expectLen:   3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d, elem := tc.setup()

			val, ok := d.Remove(elem)

			if ok != tc.expectOk {
				t.Errorf("Expected ok %v, got %v", tc.expectOk, ok)
			}
			if val != tc.expectValue {
				t.Errorf("Expected value %d, got %d", tc.expectValue, val)
			}
			if d.list.Len() != tc.expectLen {
				t.Errorf("Expected length %d, got %d", tc.expectLen, d.Len())
			}
		})
	}

	t.Run("Remove all elements sequentially", func(t *testing.T) {
		d := New[int]()
		d.PushBack(1, 2, 3)
		expected := []struct {
			val      int
			lenAfter int
		}{
			{1, 2},
			{3, 1},
			{2, 0},
		}

		for i, exp := range expected {
			var elem *list.Element
			if i == 0 {
				elem = d.list.Front() // Remove first
			} else if i == 1 {
				elem = d.list.Back() // Remove last
			} else {
				elem = d.list.Front() // Remove remaining
			}

			val, ok := d.Remove(elem)
			if !ok {
				t.Errorf("Removal %d failed", i+1)
			}
			if val != exp.val {
				t.Errorf("Expected %d, got %d in removal %d", exp.val, val, i+1)
			}
			if d.list.Len() != exp.lenAfter {
				t.Errorf("Expected length %d, got %d after removal %d", exp.lenAfter, d.Len(), i+1)
			}
		}

		if d.list.Len() != 0 {
			t.Error("Deque should be empty after all removals")
		}
	})

	// Concurrent test
	t.Run("Concurrent removal", func(t *testing.T) {
		d := New[int]()
		d.PushBack(1, 2, 3, 4, 5)
		elem := d.list.Front().Next().Next() // value 3

		done := make(chan bool)
		go func() {
			val, ok := d.Remove(elem)
			if !ok || val != 3 {
				t.Error("Concurrent removal failed")
			}
			done <- true
		}()

		// Try to access the deque while removal is happening
		for i := 0; i < 10; i++ {
			length := d.list.Len()
			if length != 5 && length != 4 {
				t.Errorf("Unexpected length during concurrent access: %d", length)
			}
		}

		<-done
		if d.list.Len() != 4 {
			t.Errorf("Expected length 4 after removal, got %d", d.Len())
		}
	})
}
