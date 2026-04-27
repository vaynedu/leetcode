package interview

import "testing"

func TestCodec(t *testing.T) {
    c := Constructor()
    root := NewTree([]any{1, 2, 3, nil, nil, 4, 5})
    serialized := c.serialize(root)
    deserialized := c.deserialize(serialized)
    if deserialized.Val != 1 || deserialized.Left.Val != 2 || deserialized.Right.Val != 3 {
        t.Errorf("deserialize failed")
    }
}
