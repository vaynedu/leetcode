package backtrack

import (
	"fmt"
	"testing"
)

func TestBacktrack(t *testing.T) {
	fmt.Println("=============================================")
	fmt.Println("       回溯算法分类模板 - 测试验证")
	fmt.Println("=============================================")
	fmt.Println()
	fmt.Println("覆盖类型:")
	fmt.Println("  1. 子集: Subsets, SubsetsWithDup")
	fmt.Println("  2. 组合: Combine, CombinationSum, CombinationSum2")
	fmt.Println("  3. 排列: Permute, PermuteUnique")
	fmt.Println("  4. 括号: GenerateParenthesis")
	fmt.Println("  5. 棋盘: SolveNQueens, TotalNQueens")
	fmt.Println("  6. 搜索: Exist")
	fmt.Println("  7. 分割: RestoreIpAddresses, Partition")
	fmt.Println()
}

// ------------------------------------------------------------
// 子集
// ------------------------------------------------------------

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	result := Subsets(nums)
	expected := [][]int{
		{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
	}
	if len(result) != len(expected) {
		t.Errorf("Subsets count = %d, want %d", len(result), len(expected))
	}
}

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	result := SubsetsWithDup(nums)
	// 去重后应该有8个子集
	if len(result) != 8 {
		t.Errorf("SubsetsWithDup count = %d, want 8", len(result))
	}
}

func TestCombine(t *testing.T) {
	result := Combine(4, 2)
	// C(4,2) = 6
	if len(result) != 6 {
		t.Errorf("Combine(4,2) count = %d, want 6", len(result))
	}
}

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := CombinationSum(candidates, target)
	// [[2,2,3],[7]]
	if len(result) != 2 {
		t.Errorf("CombinationSum count = %d, want 2", len(result))
	}
}

func TestCombinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := CombinationSum2(candidates, target)
	// 8的组合：[[1,1,6],[1,2,5],[1,7],[2,6]]
	if len(result) != 4 {
		t.Errorf("CombinationSum2 count = %d, want 4", len(result))
	}
}

// ------------------------------------------------------------
// 排列
// ------------------------------------------------------------

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	result := Permute(nums)
	// 3! = 6
	if len(result) != 6 {
		t.Errorf("Permute count = %d, want 6", len(result))
	}
}

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	result := PermuteUnique(nums)
	// 去重后3个排列
	if len(result) != 3 {
		t.Errorf("PermuteUnique count = %d, want 3", len(result))
	}
}

func TestLetterCasePermutation(t *testing.T) {
	s := "a1b2"
	result := LetterCasePermutation(s)
	// a1b2, A1b2, a1B2, A1B2
	if len(result) != 4 {
		t.Errorf("LetterCasePermutation count = %d, want 4", len(result))
	}
}

// ------------------------------------------------------------
// 括号
// ------------------------------------------------------------

func TestGenerateParenthesis(t *testing.T) {
	n := 3
	result := GenerateParenthesis(n)
	// n=3 有5种组合
	if len(result) != 5 {
		t.Errorf("GenerateParenthesis(%d) count = %d, want 5", n, len(result))
	}
}

// ------------------------------------------------------------
// N皇后
// ------------------------------------------------------------

func TestSolveNQueens(t *testing.T) {
	n := 4
	result := SolveNQueens(n)
	if len(result) != 2 {
		t.Errorf("SolveNQueens(%d) count = %d, want 2", n, len(result))
	}
}

func TestTotalNQueens(t *testing.T) {
	if n := TotalNQueens(4); n != 2 {
		t.Errorf("TotalNQueens(4) = %d, want 2", n)
	}
	if n := TotalNQueens(8); n != 92 {
		t.Errorf("TotalNQueens(8) = %d, want 92", n)
	}
}

// ------------------------------------------------------------
// 单词搜索
// ------------------------------------------------------------

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	words := []struct {
		word  string
		match bool
	}{
		{"ABCCED", true},
		{"SEE", true},
		{"ABCB", false},
	}
	for _, w := range words {
		if Exist(board, w.word) != w.match {
			t.Errorf("Exist(%s) = %v, want %v", w.word, !w.match, w.match)
		}
	}
}

// ------------------------------------------------------------
// 复原IP地址
// ------------------------------------------------------------

func TestRestoreIpAddresses(t *testing.T) {
	s := "25525511135"
	result := RestoreIpAddresses(s)
	// ["255.255.11.135","255.255.111.35"]
	if len(result) != 2 {
		t.Errorf("RestoreIpAddresses count = %d, want 2", len(result))
	}
}

// ------------------------------------------------------------
// 分割回文串
// ------------------------------------------------------------

func TestPartition(t *testing.T) {
	s := "aab"
	result := Partition(s)
	// [["a","a","b"],["aa","b"]]
	if len(result) != 2 {
		t.Errorf("Partition count = %d, want 2", len(result))
	}
}

// ------------------------------------------------------------
// 累加数
// ------------------------------------------------------------

func TestAdditiveNumber(t *testing.T) {
	tests := []struct {
		num    string
		expect bool
	}{
		{"112358", true},
		{"199100199", true},
		{"123", true},
		{"010", true},
		{"112233", false},
	}
	for _, tt := range tests {
		if AdditiveNumber(tt.num) != tt.expect {
			t.Errorf("AdditiveNumber(%s) = %v, want %v", tt.num, !tt.expect, tt.expect)
		}
	}
}

// ------------------------------------------------------------
// 预测赢家
// ------------------------------------------------------------

func TestPredictWinner(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
	}
	for _, tt := range tests {
		if PredictWinner(tt.nums) != tt.expect {
			t.Errorf("PredictWinner = %v, want %v", !tt.expect, tt.expect)
		}
	}
}
