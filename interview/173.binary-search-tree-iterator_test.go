package interview

import "testing"

func TestBSTIterator(t *testing.T) {
    root := NewTree([]any{7, 3, 15, nil, nil, 9, 20})
    iter := Constructor(root)
    vals := []int{}
    for iter.HasNext() {
        vals = append(vals, iter.Next())
    }
    want := []int{3, 7, 9, 15, 20}
    if !intEq(vals, want) {
        t.Errorf("BSTIterator got %v, want %v", vals, want)
    }
}
