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
