package interview

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("       高频面试算法题测试")
	fmt.Println("=============================================")
	fmt.Println()
}

// ================================================================
// 146. LRU 缓存
// ================================================================

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if cache.Get(1) != 1 {
		t.Errorf("Get(1) should be 1, got %d", cache.Get(1))
	}

	cache.Put(3, 3) // 淘汰 key=2

	if cache.Get(2) != -1 {
		t.Errorf("Get(2) should be -1, got %d", cache.Get(2))
	}

	cache.Put(4, 4) // 淘汰 key=1

	if cache.Get(1) != -1 {
		t.Errorf("Get(1) should be -1, got %d", cache.Get(1))
	}
	if cache.Get(3) != 3 {
		t.Errorf("Get(3) should be 3, got %d", cache.Get(3))
	}
	if cache.Get(4) != 4 {
		t.Errorf("Get(4) should be 4, got %d", cache.Get(4))
	}

	fmt.Println("✅ 146. LRUCache - PASS")
}

// ================================================================
// 56. 合并区间
// ================================================================

func TestMergeIntervals(t *testing.T) {
	tests := []struct {
		input    []Interval
		expected []Interval
	}{
		{
			[]Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[]Interval{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[]Interval{{1, 4}, {4, 5}},
			[]Interval{{1, 5}},
		},
		{
			[]Interval{{1, 4}},
			[]Interval{{1, 4}},
		},
		{
			[]Interval{},
			nil,
		},
	}

	for i, tt := range tests {
		result := merge(tt.input)
		if !intervalsEqual(result, tt.expected) {
			t.Errorf("Test %d failed", i+1)
		}
	}

	fmt.Println("✅ 56. merge - PASS")
}

// ================================================================
// 200. 岛屿数量
// ================================================================

func TestNumIslands(t *testing.T) {
	grid1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	if numIslands(grid1) != 1 {
		t.Errorf("grid1 should have 1 island")
	}

	grid2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	if numIslands(grid2) != 3 {
		t.Errorf("grid2 should have 3 islands")
	}

	grid3 := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}
	if numIslands(grid3) != 1 {
		t.Errorf("grid3 should have 1 island")
	}

	fmt.Println("✅ 200. numIslands - PASS")
}

func TestLevelOrder(t *testing.T) {
	fmt.Println("=============================================")
	fmt.Println("       102. 二叉树层序遍历")
	fmt.Println("=============================================")
	fmt.Println()

	// 构建测试树
	//     3
	//    / \
	//   9  20
	//     /  \
	//    15   7
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	result := levelOrder(root)
	expected := [][]int{{3}, {9, 20}, {15, 7}}

	if !sliceEqual(result, expected) {
		t.Errorf("levelOrder failed")
	}

	// 空树测试
	if levelOrder(nil) != nil {
		t.Errorf("nil tree should return nil")
	}

	// 单节点测试
	single := &TreeNode{Val: 1}
	if r := levelOrder(single); len(r) != 1 || r[0][0] != 1 {
		t.Errorf("single node failed")
	}

	fmt.Println("✅ 102. levelOrder - PASS")
}

func sliceEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestSummary(t *testing.T) {
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Println("       23. 合并K个升序链表")
	fmt.Println("       56. 合并区间")
	fmt.Println("       102. 二叉树层序遍历")
	fmt.Println("       146. LRU 缓存")
	fmt.Println("       200. 岛屿数量")
	fmt.Println("=============================================")
	fmt.Println("✅ All 高频面试题测试通过")
}
