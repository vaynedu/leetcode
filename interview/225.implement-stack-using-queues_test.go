package interview

import "testing"

func TestMyStack(t *testing.T) {
	s := StackConstructor()
	s.Push(1)
	s.Push(2)
	if got := s.Top(); got != 2 {
		t.Errorf("Top() = %d, want 2", got)
	}
	if got := s.Pop(); got != 2 {
		t.Errorf("Pop() = %d, want 2", got)
	}
	if got := s.Empty(); got != false {
		t.Errorf("Empty() = %v, want false", got)
	}
	s.Pop()
	if got := s.Empty(); got != true {
		t.Errorf("Empty() = %v, want true", got)
	}
}

func TestMyStackLIFOAfterMultiplePushes(t *testing.T) {
	s := StackConstructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if got := s.Pop(); got != 3 {
		t.Errorf("first Pop() = %d, want 3", got)
	}
	s.Push(4)
	for _, want := range []int{4, 2, 1} {
		if got := s.Pop(); got != want {
			t.Fatalf("Pop() = %d, want %d", got, want)
		}
	}
	if !s.Empty() {
		t.Fatal("stack should be empty after popping all elements")
	}
}
