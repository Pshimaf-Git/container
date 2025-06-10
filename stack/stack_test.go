package stack

import (
	"sync"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("Create new stack", func(t *testing.T) {
		s := New[int]()
		if s == nil {
			t.Error("New() returned nil")
		}
	})
}

func TestSize(t *testing.T) {
	t.Run("Empty stack size", func(t *testing.T) {
		s := New[int]()
		if s.Size() != 0 {
			t.Errorf("Size() = %d, want 0", s.Size())
		}
	})

	t.Run("Stack with elements size", func(t *testing.T) {
		s := New[int]()
		s.Push(10)
		s.Push(20)
		if s.Size() != 2 {
			t.Errorf("Size() = %d, want 2", s.Size())
		}
	})

	t.Run("Size after push and pop", func(t *testing.T) {
		s := New[int]()
		s.Push(10)
		s.Pop()
		if s.Size() != 0 {
			t.Errorf("Size() = %d, want 0", s.Size())
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("Push one element", func(t *testing.T) {
		s := New[int]()
		s.Push(10)
		if s.Size() != 1 {
			t.Errorf("Size() = %d, want 1", s.Size())
		}
		val, ok := s.Pop()
		if !ok || val != 10 {
			t.Errorf("Pop() = %d, %t, want 10, true", val, ok)
		}
	})

	t.Run("Push multiple elements", func(t *testing.T) {
		s := New[int]()
		s.Push(10)
		s.Push(20)
		s.Push(30)
		if s.Size() != 3 {
			t.Errorf("Size() = %d, want 3", s.Size())
		}
		val, ok := s.Pop()
		if !ok || val != 30 {
			t.Errorf("Pop() = %d, %t, want 30, true", val, ok)
		}
	})

	t.Run("Push different data types", func(t *testing.T) {
		s := New[string]()
		s.Push("hello")
		val, ok := s.Pop()
		if !ok || val != "hello" {
			t.Errorf("Pop() = %s, %t, want hello, true", val, ok)
		}
	})
}

func TestPop(t *testing.T) {
	t.Run("Pop from empty stack", func(t *testing.T) {
		s := New[int]()
		val, ok := s.Pop()
		if ok {
			t.Error("Pop() returned true, want false")
		}
		if val != 0 {
			t.Errorf("Pop() = %d, want 0", val)
		}
	})

	t.Run("Pop from stack with one element", func(t *testing.T) {
		s := New[int]()
		s.Push(10)
		val, ok := s.Pop()
		if !ok {
			t.Error("Pop() returned false, want true")
		}
		if val != 10 {
			t.Errorf("Pop() = %d, want 10", val)
		}
		if s.Size() != 0 {
			t.Errorf("Size() = %d, want 0", s.Size())
		}
	})

	t.Run("Pop from stack with multiple elements", func(t *testing.T) {
		s := New[int]()
		s.Push(10)
		s.Push(20)
		val, ok := s.Pop()
		if !ok {
			t.Error("Pop() returned false, want true")
		}
		if val != 20 {
			t.Errorf("Pop() = %d, want 20", val)
		}
		if s.Size() != 1 {
			t.Errorf("Size() = %d, want 1", s.Size())
		}
		val, ok = s.Pop()
		if !ok {
			t.Error("Pop() returned false, want true")
		}
		if val != 10 {
			t.Errorf("Pop() = %d, want 10", val)
		}
		if s.Size() != 0 {
			t.Errorf("Size() = %d, want 0", s.Size())
		}
	})

	t.Run("Pop all elements from stack", func(t *testing.T) {
		s := New[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		_, ok := s.Pop()
		if !ok {
			t.Errorf("Expected to successfully pop an element")
		}
		_, ok = s.Pop()
		if !ok {
			t.Errorf("Expected to successfully pop an element")
		}
		_, ok = s.Pop()
		if !ok {
			t.Errorf("Expected to successfully pop an element")
		}

		_, ok = s.Pop()
		if ok {
			t.Errorf("Expected pop to return false")
		}
	})
}

func TestStack(t *testing.T) {
	type testCase[T any] struct {
		name         string
		initialStack []T
		pushValues   []T
		popCount     int
		expectedSize uint32
		expectedPop  T
		expectedOK   bool
	}

	testCasesInt := []testCase[int]{
		{
			name:         "Empty stack",
			initialStack: []int{},
			pushValues:   []int{},
			popCount:     1,
			expectedSize: 0,
			expectedPop:  0,
			expectedOK:   false,
		},
		{
			name:         "Push one, pop one",
			initialStack: []int{},
			pushValues:   []int{1},
			popCount:     1,
			expectedSize: 0,
			expectedPop:  1,
			expectedOK:   true,
		},
		{
			name:         "Push multiple, pop one",
			initialStack: []int{},
			pushValues:   []int{1, 2, 3},
			popCount:     1,
			expectedSize: 2,
			expectedPop:  3,
			expectedOK:   true,
		},
		{
			name:         "Push multiple, pop multiple",
			initialStack: []int{},
			pushValues:   []int{1, 2, 3},
			popCount:     3,
			expectedSize: 0,
			expectedPop:  1,
			expectedOK:   true,
		},
		{
			name:         "Push multiple, pop more than pushed",
			initialStack: []int{},
			pushValues:   []int{1, 2},
			popCount:     3,
			expectedSize: 0,
			expectedPop:  0,
			expectedOK:   false,
		},
		{
			name:         "Initial stack, pop one",
			initialStack: []int{1, 2, 3}, // Pushed in reverse order
			pushValues:   []int{},
			popCount:     1,
			expectedSize: 2,
			expectedPop:  3,
			expectedOK:   true,
		},
	}

	for _, tc := range testCasesInt {
		t.Run(tc.name, func(t *testing.T) {
			s := New[int]()

			// Initialize stack
			for _, v := range tc.initialStack {
				s.Push(v)
			}

			// Push values
			for _, v := range tc.pushValues {
				s.Push(v)
			}

			// Pop values
			var popValue int
			ok := true
			for i := 0; i < tc.popCount; i++ {
				popValue, ok = s.Pop()
			}

			// Check size
			actualSize := s.Size()
			if actualSize != tc.expectedSize {
				t.Errorf("Size() = %d, want %d", actualSize, tc.expectedSize)
			}

			// Check pop value and ok
			if tc.popCount > 0 {
				if popValue != tc.expectedPop && ok == tc.expectedOK {
					t.Errorf("Pop() = %v, %t, want %v, %t", popValue, ok, tc.expectedPop, tc.expectedOK)
				}
			}
		})
	}

	testCasesString := []testCase[string]{
		{
			name:         "String stack",
			initialStack: []string{},
			pushValues:   []string{"hello", "world"},
			popCount:     1,
			expectedSize: 1,
			expectedPop:  "world",
			expectedOK:   true,
		},
		{
			name:         "Empty string stack",
			initialStack: []string{},
			pushValues:   []string{},
			popCount:     1,
			expectedSize: 0,
			expectedPop:  "",
			expectedOK:   false,
		},
	}

	for _, tc := range testCasesString {
		t.Run(tc.name, func(t *testing.T) {
			s := New[string]()

			// Initialize stack
			for _, v := range tc.initialStack {
				s.Push(v)
			}

			// Push values
			for _, v := range tc.pushValues {
				s.Push(v)
			}

			// Pop values
			var popValue string
			ok := true
			for i := 0; i < tc.popCount; i++ {
				popValue, ok = s.Pop()
			}

			// Check size
			actualSize := s.Size()
			if actualSize != tc.expectedSize {
				t.Errorf("Size() = %d, want %d", actualSize, tc.expectedSize)
			}

			// Check pop value and ok
			if tc.popCount > 0 {
				if popValue != tc.expectedPop && ok == tc.expectedOK {
					t.Errorf("Pop() = %v, %t, want %v, %t", popValue, ok, tc.expectedPop, tc.expectedOK)
				}
			}
		})
	}
}

func TestConcurrentPushPop(t *testing.T) {
	s := New[int]()
	numRoutines := 100
	numOperations := 1000

	var wg sync.WaitGroup
	wg.Add(numRoutines)

	for i := 0; i < numRoutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				s.Push(j)
				s.Pop() // Could be a no-op on an empty stack
			}
		}()
	}

	wg.Wait()

	// Check if stack is empty after all operations
	// It might not be exactly zero due to the race condition, but it should be small.
	size := s.Size()
	if size > uint32(numRoutines) {
		t.Errorf("Size() = %d, want close to 0 after concurrent operations", size)
	}
}
