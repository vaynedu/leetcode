package leetcode

// 113.路径总和 II

import (
	"testing"
)

/**
给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。

 叶子节点 是指没有子节点的节点。







 示例 1：


输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]


 示例 2：


输入：root = [1,2,3], targetSum = 5
输出：[]


 示例 3：


输入：root = [1,2], targetSum = 0
输出：[]




 提示：


 树中节点总数在范围 [0, 5000] 内
 -1000 <= Node.val <= 1000
 -1000 <= targetSum <= 1000


 Related Topics 树 深度优先搜索 回溯 二叉树 👍 1177 👎 0

*/

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func BackTrack(root *TreeNode, targetSum int, curSum int, path []int, res *[][]int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if curSum+root.Val == targetSum {
			// 深拷贝 path 切片
			temp := make([]int, len(path))
			copy(temp, path)
			*res = append(*res, temp)
		}
		return
	}

	BackTrack(root.Left, targetSum, curSum+root.Val, path, res)
	BackTrack(root.Right, targetSum, curSum+root.Val, path, res)

	// 回溯路径
	path = path[:len(path)-1]
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	BackTrack(root, targetSum, 0, []int{}, &ans)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func TestPathSumIi(t *testing.T) {
	tests := []struct {
		name      string
		root      *TreeNode
		targetSum int
		want      [][]int
	}{
		{
			name: "Example 1",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{
						Val: 13,
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 1,
						},
					},
				},
			},
			targetSum: 22,
			want:      [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "Example 2",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			targetSum: 5,
			want:      [][]int{},
		},
		{
			name: "Example 3",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			targetSum: 0,
			want:      [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathSum(tt.root, tt.targetSum)
			if !equal(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// equal 比较两个二维整数切片是否相等
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !intSliceEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// intSliceEqual 比较两个整数切片是否相等
func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
