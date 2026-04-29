package interview

import "testing"

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if got := s.GetMin(); got != -3 {
		t.Errorf("GetMin() = %d, want -3", got)
	}
	s.Pop()
	if got := s.Top(); got != 0 {
		t.Errorf("Top() = %d, want 0", got)
	}
	if got := s.GetMin(); got != -2 {
		t.Errorf("GetMin() = %d, want -2", got)
	}
}

func TestMinStackDuplicateMinimum(t *testing.T) {
	s := NewMinStack()
	s.Push(2)
	s.Push(1)
	s.Push(1)
	s.Pop()
	if got := s.GetMin(); got != 1 {
		t.Errorf("GetMin() after popping duplicate minimum = %d, want 1", got)
	}
	s.Pop()
	if got := s.GetMin(); got != 2 {
		t.Errorf("GetMin() after popping all minimums = %d, want 2", got)
	}
}
