package backtrack

// ============================================================================
// 回溯算法（Backtracking）分类模板
// 适用场景：组合、排列、切割、子集、棋盘、N皇后、括号生成等
// 核心思想：选/不选，选了之后递归，撤销选择
// ============================================================================

import (
	"fmt"
)

// ============================================================================
// 一、子集（Subset）
// 特征：每个元素选或不选，形成所有子集
// 模板：for 遍历元素，收集路径
// ============================================================================

// Subsets 78. 子集
// 每个元素都可以选或不选，共 2^n 个子集
func Subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		// 收集结果（所有节点都要收集）
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		// 从 idx 开始选
		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1] // 撤销选择
		}
	}

	dfs(0)
	return res
}

// SubsetsWithDup 90. 子集 II（包含重复元素）
// SubsetsWithDup 90. 子集 II（包含重复元素）
// 先排序，再剪枝（跳过相同元素）
func SubsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	// 先排序，让相同元素相邻
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	// 简单的插入排序
	for i := 1; i < len(sorted); i++ {
		j := i
		for j > 0 && sorted[j] < sorted[j-1] {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
	}

	var dfs func(idx int)
	dfs = func(idx int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		seen := make(map[int]bool)
		for i := idx; i < len(sorted); i++ {
			if seen[sorted[i]] {
				continue
			}
			seen[sorted[i]] = true
			path = append(path, sorted[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}

// ============================================================================
// 二、组合（Combination）
// 特征：从 n 个数中选 k 个数的所有组合
// 模板：for 遍历，递归选，撤销
// ============================================================================

// Combine 77. 组合
// 从 1..n 中选 k 个数的组合
func Combine(n, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(start int)
	dfs = func(start int) {
		// 终止条件：选够 k 个
		if len(path) == k {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		// 剪枝：不足够剩余元素时跳过
		// 还需选 (k - len(path)) 个，剩余可用 [start..n]
		// 所以 start 最大为 n - (k-len(path)) + 1
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return res
}

// CombinationSum 39. 组合总和（无重复元素，每个元素可重复选）
// candidates 无重复，正整数数组
func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(idx int, remain int)
	dfs = func(idx int, remain int) {
		if remain < 0 {
			return
		}
		if remain == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, remain-candidates[i]) // 注意：i 不变，允许重复选
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)
	return res
}

// CombinationSum2 40. 组合总和 II（每个数字只能用一次，有重复）
// candidates 有重复数字，每个数字只能用一次
func CombinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	// 先排序
	sorted := make([]int, len(candidates))
	copy(sorted, candidates)
	for i := 1; i < len(sorted); i++ {
		j := i
		for j > 0 && sorted[j] < sorted[j-1] {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
	}

	used := make([]bool, len(sorted))

	var dfs func(idx int, remain int)
	dfs = func(idx int, remain int) {
		if remain < 0 {
			return
		}
		if remain == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := idx; i < len(sorted); i++ {
			// 剪枝：跳过同一层相同值
			if i > 0 && sorted[i] == sorted[i-1] && !used[i-1] {
				continue
			}
			path = append(path, sorted[i])
			used[i] = true
			dfs(i+1, remain-sorted[i])
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs(0, target)
	return res
}

// ============================================================================
// 三、排列（Permutation）
// 特征：所有元素的顺序都算一种排列
// 模板：用 used 数组标记已选，遍历所有元素
// ============================================================================

// Permute 46. 全排列
// 无重复数字的全排列
func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs()
	return res
}

// PermuteUnique 47. 全排列 II（有重复数字）
func PermuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))

	// 先排序
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	for i := 1; i < len(sorted); i++ {
		j := i
		for j > 0 && sorted[j] < sorted[j-1] {
			sorted[j], sorted[j-1] = sorted[j-1], sorted[j]
			j--
		}
	}

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(sorted); i++ {
			// 剪枝：同层相同值跳过
			if used[i] {
				continue
			}
			if i > 0 && sorted[i] == sorted[i-1] && !used[i-1] {
				continue
			}
			path = append(path, sorted[i])
			used[i] = true
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs()
	return res
}

// ============================================================================
// 四、排列式子（LetterCasePermutation）
// 特征：数字字符串中，数字不变，字母变大或小写
// ============================================================================

// LetterCasePermutation 784. 字母大小写全排列
func LetterCasePermutation(s string) []string {
	res := make([]string, 0)
	path := make([]byte, len(s))

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			res = append(res, string(path))
			return
		}

		path[idx] = s[idx]
		dfs(idx + 1)

		// 如果是字母，则尝试转换大小写
		c := s[idx]
		if c >= 'A' && c <= 'Z' {
			path[idx] = c + 32 // 转小写
			dfs(idx + 1)
		} else if c >= 'a' && c <= 'z' {
			path[idx] = c - 32 // 转大写
			dfs(idx + 1)
		}
	}

	dfs(0)
	return res
}

// ============================================================================
// 五、括号生成（Parentheses Generation）
// 特征：生成 n 对括号的所有合法组合
// 模板：左右括号计数，选的条件是左<右
// ============================================================================

// GenerateParenthesis 22. 括号生成
// n 对括号的所有合法组合
func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0)

	var dfs func(left, right int)
	dfs = func(left, right int) {
		// 终止条件：左右都用完
		if left == n && right == n {
			res = append(res, string(path))
			return
		}

		// 可以加左括号（只要左括号没用完）
		if left < n {
			path = append(path, '(')
			dfs(left+1, right)
			path = path[:len(path)-1]
		}

		// 可以加右括号（只要右括号 < 左括号）
		if right < left {
			path = append(path, ')')
			dfs(left, right+1)
			path = path[:len(path)-1]
		}
	}

	dfs(0, 0)
	return res
}

// ============================================================================
// 六、棋盘类（N-Queens, Sudoku, Word Search）
// 特征：在二维网格上搜索，需要路径回溯
// ============================================================================

// SolveNQueens 51. N 皇后
// n 皇后问题：在 n×n 棋盘上放 n 个皇后，互相不攻击
// 解法：按行放置，用 cols, diag1, diag2 标记列和两条对角线
func SolveNQueens(n int) [][]string {
	res := make([][]string, 0)
	path := make([][]byte, n)

	// 初始化棋盘
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		path[i] = row
	}

	// 用于按行递归，col记录该行的皇后列
	var dfs func(row int, cols map[int]bool, diag1 map[int]bool, diag2 map[int]bool)
	dfs = func(row int, cols map[int]bool, diag1 map[int]bool, diag2 map[int]bool) {
		if row == n {
			// 收集结果
			sol := make([]string, n)
			for i := 0; i < n; i++ {
				sol[i] = string(path[i])
			}
			res = append(res, sol)
			return
		}

		for col := 0; col < n; col++ {
			d1 := row - col + n // 主对角线 row-col 相同
			d2 := row + col     // 副对角线 row+col 相同

			if cols[col] || diag1[d1] || diag2[d2] {
				continue // 被攻击，跳过
			}

			// 放置皇后
			path[row][col] = 'Q'
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			dfs(row+1, cols, diag1, diag2)

			// 撤销
			path[row][col] = '.'
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	cols := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	dfs(0, cols, diag1, diag2)
	return res
}

// TotalNQueens 52. N 皇后 II
// 返回 n 皇后问题的解法数量
func TotalNQueens(n int) int {
	count := 0

	var dfs func(row int, cols map[int]bool, diag1 map[int]bool, diag2 map[int]bool)
	dfs = func(row int, cols map[int]bool, diag1 map[int]bool, diag2 map[int]bool) {
		if row == n {
			count++
			return
		}

		for col := 0; col < n; col++ {
			d1 := row - col + n
			d2 := row + col

			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			dfs(row+1, cols, diag1, diag2)

			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	cols := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	dfs(0, cols, diag1, diag2)
	return count
}

// Exist 79. 单词搜索
// 在二维网格中搜索单词，格子不能重复使用
func Exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	// 方向：上下左右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	var dfs func(i, j, idx int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word) {
			return true
		}

		// 越界或已访问或字符不匹配
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || board[i][j] != word[idx] {
			return false
		}

		visited[i][j] = true

		for _, d := range dirs {
			if dfs(i+d[0], j+d[1], idx+1) {
				visited[i][j] = false
				return true
			}
		}

		visited[i][j] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// ============================================================================
// 七、路径恢复类（RestoreIPAddresses, IP白板）
// 特征：将字符串分割成合法格式
// ============================================================================

// RestoreIpAddresses 93. 复原 IP 地址
// 将字符串分割成有效 IP 地址（0.0.0.0 ~ 255.255.255.255）
func RestoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		// 已经分成4段
		if len(path) == 4 {
			if idx == len(s) {
				res = append(res, joinStrings(path))
			}
			return
		}

		// 还没分完但已经用完字符串
		if idx == len(s) {
			return
		}

		// 剪枝：剩余字符太多或太少
		remain := len(s) - idx
		needMin := 4 - len(path)       // 最少还需needMin段，每段至少1个
		needMax := (4 - len(path)) * 3 // 最多还需needMax段，每段最多3个
		if remain < needMin || remain > needMax {
			return
		}

		// 尝试取1-3个字符作为一段
		for length := 1; length <= 3 && idx+length <= len(s); length++ {
			seg := s[idx : idx+length]
			// 检查合法性：不能有前导0（0本身合法，01/001不合法）
			if length > 1 && seg[0] == '0' {
				continue
			}
			// 检查数值 <= 255
			if length == 3 {
				val := (seg[0]-'0')*100 + (seg[1]-'0')*10 + (seg[2] - '0')
				if val > 255 {
					continue
				}
			}
			path = append(path, seg)
			dfs(idx + length)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}

func joinStrings(strs []string) string {
	result := strs[0]
	for i := 1; i < len(strs); i++ {
		result += "." + strs[i]
	}
	return result
}

// ============================================================================
// 八、分割类（PalindromePartitioning）
// 特征：将字符串分割成若干子串，每个子串满足某种条件
// ============================================================================

// Partition 131. 分割回文串
// 将字符串分割成若干子串，每个子串都是回文
func Partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)

	isPalindrome := func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := idx; i < len(s); i++ {
			if isPalindrome(idx, i) {
				path = append(path, s[idx:i+1])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return res
}

// ============================================================================
// 九、组合/排列计数
// ============================================================================

// AdditiveNumber 306. 累加数
// 判断字符串是否是累加序列（如 112358）
func AdditiveNumber(num string) bool {
	n := len(num)

	// 从 i 和 j 位置开始取前两个数
	var dfs func(idx int, prev []int) bool
	dfs = func(idx int, prev []int) bool {
		if idx == n {
			return len(prev) >= 3 // 至少需要两个加法
		}

		// 尝试取下一个数
		for length := 1; idx+length <= n; length++ {
			if length > 1 && num[idx] == '0' { // 不能有前导0
				break
			}
			curr := parseInt(num[idx : idx+length])
			if len(prev) >= 2 {
				last := prev[len(prev)-1]
				secondLast := prev[len(prev)-2]
				if int64(last)+int64(secondLast) != int64(curr) {
					continue
				}
			}
			prev = append(prev, curr)
			if dfs(idx+length, prev) {
				return true
			}
			prev = prev[:len(prev)-1]
		}
		return false
	}

	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if num[0] == '0' && i > 1 {
				break
			}
			if num[i] == '0' && j-i > 1 {
				continue
			}
			first := parseInt(num[0:i])
			second := parseInt(num[i:j])
			if dfs(j, []int{first, second}) {
				return true
			}
		}
	}
	return false
}

func parseInt(s string) int {
	val := 0
	for i := 0; i < len(s); i++ {
		val = val*10 + int(s[i]-'0')
	}
	return val
}

// ============================================================================
// 十、比赛博弈类（CanWin, PredictWinner）
// ============================================================================

// PredictWinner 486. 预测赢家
// 先手能否赢：每次从数组两端选一个，返回差值
func PredictWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	// dp[i][j] = 玩家在 [i,j] 能拿到的最优分数（减去对手的）
	for length := 2; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			// 选左边或选右边，然后对手在剩余区间最优
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] >= 0
}

// ============================================================================
// 辅助函数
// ============================================================================

// PrintResults 打印结果（用于调试）
func PrintResults(res interface{}) {
	fmt.Printf("%v\n", res)
}
