// package deque is a linked list based implementation of a double

// ended queue

package deque

import (
	"reflect"
	"testing"
)

func Test_revers(t *testing.T) {
	tests := []struct {
		name string
		a    []any
		want []any
	}{
		{
			name: "base case",
			a:    []any{1, 2, 3},
			want: []any{3, 2, 1},
		},

		{
			name: "empty input slice",
			a:    []any{},
			want: []any{},
		},

		{
			name: "nil input slice",
			a:    nil,
			want: nil,
		},

		{
			name: "\"palindrom\" input slice",
			a:    []any{3, 0, 0, 3},
			want: []any{3, 0, 0, 3},
		},

		{
			name: "base-case",
			a:    []any{1, '2', "3"},
			want: []any{"3", '2', 1},
		},

		{
			name: "mixed input slice",
			a:    []any{3, "hello", complex64(3.43 * 23), uint(34), nil, '1'},
			want: []any{'1', nil, uint(34), complex64(3.43 * 23), "hello", 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revers(tt.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_Len(t *testing.T) {
	fullDeque := New()
	fullDeque.PushFront(1, 2, 3, 4)

	tests := []struct {
		name string
		d    *Deque
		want int
	}{
		{
			name: "empty deque",
			d:    New(),
			want: 0,
		},

		{
			name: "full deque",
			d:    fullDeque,
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Len(); got != tt.want {
				t.Errorf("Deque.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_IsEmpty(t *testing.T) {
	d := New()
	d.PushFront(1)

	tests := []struct {
		name string
		d    *Deque
		want bool
	}{
		{
			name: "empty deque",
			d:    New(),
			want: true,
		},

		{
			name: "full deque",
			d:    d,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.IsEmpty(); got != tt.want {
				t.Errorf("Deque.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_ToArray(t *testing.T) {
	mixed := New()
	mixedValues := []any{1, uint(2), 0.57, 4, 5 * 2i, false, 7, true, "9", int16(10)}
	mixed.PushBack(mixedValues...)

	basic := New()
	values := []any{1, "hello"}
	basic.PushBack(values...)

	tests := []struct {
		name string
		d    *Deque
		want []any
	}{
		{
			name: "empty deque",
			d:    New(),
			want: []any{},
		},

		{
			name: "mixed types",
			d:    mixed,
			want: mixedValues,
		},

		{
			name: "base-case",
			d:    basic,
			want: values,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.ToArray(); !reflect.DeepEqual(got, tt.want) {
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
		name string
		d    *Deque
		args args
		want []any
	}{
		{
			name: "push one",
			d:    New(),
			args: args{
				values: []any{1},
			},
			want: []any{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.PushFront(tt.args.values...)
			got := tt.d.ToArray()

			if !reflect.DeepEqual(got, tt.want) {
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
		name string
		d    *Deque
		args args
		want []any
	}{
		{
			name: "push one",
			d:    New(),
			args: args{
				values: []any{1},
			},
			want: []any{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.PushFront(tt.args.values...)
			got := tt.d.ToArray()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.PushBack(%v) got = %v, want = %v", tt.args.values, got, tt.want)
			}
		})
	}
}

func TestDeque_Clear(t *testing.T) {
	fullDeque := New()
	values := []any{1, true, 5 * 2i, uintptr(3), "hello", ' '}
	fullDeque.PushBack(values...)

	tests := []struct {
		name string
		d    *Deque
		want int
	}{
		{
			name: "empty deque1",
			d:    New(),
			want: 0,
		},

		{
			name: "full deque",
			d:    fullDeque,
			want: len(values),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Clear(); got != tt.want || tt.d.Len() != 0 {
				t.Errorf("Deque.Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeque_PopFront(t *testing.T) {
	fullDeque := New()
	values := []any{nil, true, 5 * 2i, uintptr(3), "hello", ' '}
	fullDeque.PushBack(values...)

	tests := []struct {
		name  string
		d     *Deque
		want  any
		want1 bool
	}{
		{
			name:  "empty deque",
			d:     New(),
			want:  nil,
			want1: false,
		},

		{
			name:  "full deque",
			d:     fullDeque,
			want:  values[0],
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.PopFront()
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
	fullDeque := New()
	values := []any{nil, true, 5 * 2i, uintptr(3), "hello", ' '}
	fullDeque.PushBack(values...)

	tests := []struct {
		name  string
		d     *Deque
		want  any
		want1 bool
	}{
		{
			name:  "empty deque",
			d:     New(),
			want:  nil,
			want1: false,
		},

		{
			name:  "full deque",
			d:     fullDeque,
			want:  ' ',
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.PopBack()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deque.PopBack() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Deque.PopBack() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
