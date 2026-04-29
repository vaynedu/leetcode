package interview

import "testing"

func TestLevelOrderSingleProblemCases(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want [][]int
	}{
		{
			name: "基本三层",
			root: NewTree([]any{3, 9, 20, nil, nil, 15, 7}),
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name: "单节点",
			root: NewTree([]any{1}),
			want: [][]int{{1}},
		},
		{
			name: "不完整二叉树",
			root: NewTree([]any{1, 2, 3, nil, 4, nil, 5}),
			want: [][]int{{1}, {2, 3}, {4, 5}},
		},
		{
			name: "空树",
			root: nil,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := levelOrder(tt.root)
			if !intSliceEq(got, tt.want) {
				t.Fatalf("levelOrder() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
