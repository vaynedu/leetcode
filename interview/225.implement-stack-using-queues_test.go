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
