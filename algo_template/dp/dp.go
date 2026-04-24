package dp

// ============================================================================
// 动态规划分类模板
// ============================================================================
// 类型：坐标型 | 序列型 | 划分型 | 背包 | 区间型 | 状态压缩 | 树形
// ============================================================================

import "math"

// ============================================================================
// 一、坐标型 DP
// 特征：网格/数组上移动，dp[i][j] 表示走到 (i,j) 的最优解
// 模板：dp[i][j] = f(dp[i-1][j], dp[i][j-1])
// ============================================================================

// MinPathSum 64. 最小路径和
// 只能向下或向右走，求从左上到右下的最小路径和
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化
	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 填表
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}

// UniquePaths 62. 不同路径
// 只能向下或向右走，求不同路径数
func UniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// UniquePathsWithObstacles 63. 不同路径 II（有障碍物）
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// CalculateMinimumHP 174. 地下城游戏（从后向前）
// 战士从左上到右下，最少需要多少初始血量才能存活
func CalculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 {
		return 0
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[m][n-1] = 1
	dp[m-1][n] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if need < 1 {
				need = 1
			}
			dp[i][j] = need
		}
	}
	return dp[0][0]
}

// ============================================================================
// 二、序列型 DP
// 特征：dp[i] 表示以第 i 个元素结尾/截至 i 的最优解
// 模板：dp[i] = max(dp[j] + ...) (j < i, 满足条件)
// ============================================================================

// LengthOfLIS 300. 最长上升子序列
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

// Rob 198. 打家劫舍（序列型，依赖前两个）
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}

// NumDecodings 91. 解码方法
func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		// 单字符解码
		if s[i-1] != '0' {
			dp[i] = dp[i-1]
		}
		// 双字符解码
		two := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if two >= 10 && two <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}

// WordBreak 139. 单词划分
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, w := range wordDict {
		wordSet[w] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}

// ============================================================================
// 三、划分型 DP
// 特征：将序列划分成若干段，取最值/计数
// 模板：dp[i] = min(dp[j] + 1) (j < i, 满足条件)
// ============================================================================

// MinCut 132. 分割回文串 II
// 返回最少分割次数，使每个子串都是回文
func MinCut(s string) int {
	n := len(s)
	// 预处理：isPalin[i][j] 表示 s[i:j+1] 是否是回文
	isPalin := make([][]bool, n)
	for i := range isPalin {
		isPalin[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			isPalin[i][j] = s[i] == s[j] && (j-i < 2 || isPalin[i+1][j-1])
		}
	}

	// dp[i] = s[0:i+1] 的最少分割次数
	dp := make([]int, n)
	for i := range dp {
		if isPalin[0][i] {
			dp[i] = 0
		} else {
			dp[i] = i
			for j := 0; j < i; j++ {
				if isPalin[j+1][i] {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}
	return dp[n-1]
}

// SplitArray 410. 分割数组的最大值（划分成 m 个子数组，最小化最大值）
func SplitArray(nums []int, m int) int {
	n := len(nums)
	// 前缀和
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// dp[i][k] = 以 nums[i] 结尾，分成 k 组的最小最大和
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

	// 初始化：分成1组
	for i := 0; i < n; i++ {
		dp[i][1] = prefix[i+1]
	}

	for k := 2; k <= m; k++ {
		for i := k - 1; i < n; i++ {
			// 枚举上一次分割点 j
			for j := k - 2; j < i; j++ {
				cur := max(dp[j][k-1], prefix[i+1]-prefix[j+1])
				dp[i][k] = min(dp[i][k], cur)
			}
		}
	}
	return dp[n-1][m]
}

// ============================================================================
// 四、背包 DP
// 特征：选或不选，容量限制，极值
// 01背包：j 从大到小 | 完全背包：j 从小到大
// ============================================================================

// CanPartition 416. 分割等和子集（01背包）
func CanPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, v := range nums {
		for j := target; j >= v; j-- {
			dp[j] = dp[j] || dp[j-v]
		}
	}
	return dp[target]
}

// FindTargetSumWays 494. 目标和（01背包计数）
func FindTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if abs(target) > sum {
		return 0
	}
	// 转化为：选一部分凑 (sum+target)/2
	cap := (sum + target) / 2
	if cap < 0 {
		cap = -cap
	}

	dp := make([]int, cap+1)
	dp[0] = 1
	for _, v := range nums {
		for j := cap; j >= v; j-- {
			dp[j] += dp[j-v]
		}
	}
	return dp[cap]
}

// CoinChange 322. 零钱兑换（完全背包）
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for _, c := range coins {
		for j := c; j <= amount; j++ {
			if dp[j-c] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// FindMaxForm 474. 一和零（二维费用01背包）
func FindMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zero, one := count01(s)
		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zero][j-one]+1)
			}
		}
	}
	return dp[m][n]
}

func count01(s string) (int, int) {
	zero, one := 0, 0
	for _, c := range s {
		if c == '0' {
			zero++
		} else {
			one++
		}
	}
	return zero, one
}

// ============================================================================
// 五、区间型 DP
// 特征：处理区间 [i,j]，按长度递增递推
// 模板：dp[i][j] = min(dp[i][k] + dp[k+1][j] + cost)
// ============================================================================

// LongestPalindromeSubseq 516. 最长回文子序列
func LongestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	// 按长度递增
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if length == 2 {
					dp[i][j] = 2
				} else {
					dp[i][j] = dp[i+1][j-1] + 2
				}
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

// BurstBalloons 312. 戳气球（区间划分+边界）
func BurstBalloons(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 扩展数组，两端加1
	arr := make([]int, n+2)
	arr[0] = 1
	arr[n+1] = 1
	for i := 0; i < n; i++ {
		arr[i+1] = nums[i]
	}

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// 长度从3到n+2（对应戳1到n个气球）
	for length := 3; length <= n+2; length++ {
		for i := 0; i+length-1 <= n+1; i++ {
			j := i + length - 1
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+arr[i]*arr[k]*arr[j])
			}
		}
	}
	return dp[0][n+1]
}

// StoneGame 877. 石子游戏（区间博弈）
func StoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}

	for d := 1; d < n; d++ {
		for i := 0; i+d < n; i++ {
			j := i + d
			dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}

// ============================================================================
// 六、状态压缩 DP
// 特征：用二进制 mask 表示集合状态
// 模板：dp[mask][i] = min(dp[mask_without_i][j] + dist[j][i])
// ============================================================================

// CanWin 464. 我能赢吗（博弈+状压）
func CanWin(maxChoosableInt int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return true
	}
	if maxChoosableInt*(maxChoosableInt+1)/2 < desiredTotal {
		return false
	}

	used := 0 // bitmask of used numbers
	memo := make(map[int]bool)

	var dfs func(used int, cur int) bool
	dfs = func(used int, cur int) bool {
		if v, ok := memo[used]; ok {
			return v
		}
		for i := 1; i <= maxChoosableInt; i++ {
			bit := 1 << (i - 1)
			if used&bit == 0 { // i 未被选
				// 选 i，如果能赢或让对手输
				next := used | bit
				if cur+i >= desiredTotal || !dfs(next, cur+i) {
					memo[used] = true
					return true
				}
			}
		}
		memo[used] = false
		return false
	}

	return dfs(used, 0)
}

// CanPartitionKSubsets 698. 划分为k个相等子集
func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	n := len(nums)

	// 从大到小排序，剪枝
	sortDesc(nums)

	if nums[0] > target {
		return false
	}

	size := 1 << n
	dp := make([]int, size)
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for mask := 0; mask < size; mask++ {
		if dp[mask] == -1 {
			continue
		}
		// 找到一个未填满的桶，装 nums[i]
		for i := 0; i < n; i++ {
			next := mask | (1 << i)
			if mask&(1<<i) == 0 && dp[mask]+nums[i] <= target && dp[next] == -1 {
				dp[next] = (dp[mask] + nums[i]) % target
			}
		}
	}
	return dp[size-1] == 0
}

func sortDesc(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

// TSP 旅行商问题（状态压缩）
func TSP(dist [][]int) int {
	n := len(dist)
	size := 1 << n
	dp := make([][]int, size)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	dp[1][0] = 0 // 起点0，已访问集合{0}

	for mask := 0; mask < size; mask++ {
		for i := 0; i < n; i++ {
			if dp[mask][i] == math.MaxInt {
				continue
			}
			for j := 0; j < n; j++ {
				if mask&(1<<j) == 0 { // j未访问
					next := mask | (1 << j)
					dp[next][j] = min(dp[next][j], dp[mask][i]+dist[i][j])
				}
			}
		}
	}

	fullMask := size - 1
	result := math.MaxInt
	for i := 1; i < n; i++ {
		result = min(result, dp[fullMask][i]+dist[i][0])
	}
	return result
}

// ============================================================================
// 七、树形 DP
// 特征：在树上做 DP，通常用后序遍历
// 模板：父节点 = f(左子树, 右子树)
// ============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DiameterOfBinaryTree 543. 二叉树直径
// 任意两个节点之间最长路径（不一定过根节点）
func DiameterOfBinaryTree(root *TreeNode) int {
	result := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		result = max(result, left+right) // 经过当前节点的路径
		return max(left, right) + 1      // 返回高度
	}

	dfs(root)
	return result
}

// MaxPathSum 124. 二叉树中的最大路径和
func MaxPathSum(root *TreeNode) int {
	result := math.MinInt

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		// 四种情况：只自己、+左、+右、+左右
		curMax := node.Val
		if left > 0 {
			curMax += left
		}
		if right > 0 {
			curMax += right
		}
		result = max(result, curMax)

		// 返回给父节点：只能选一边或自己
		return max(node.Val, max(node.Val+left, node.Val+right))
	}

	dfs(root)
	return result
}

// RobIII 337. 打家劫舍 III（树形DP + 选/不选）
func RobIII(root *TreeNode) int {
	result := robDFS(root)
	return max(result[0], result[1])
}

func robDFS(node *TreeNode) [2]int {
	if node == nil {
		return [2]int{0, 0}
	}
	left := robDFS(node.Left)
	right := robDFS(node.Right)

	// 不选当前节点：取两个孩子选/不选的最大和
	notSelect := max(left[0], left[1]) + max(right[0], right[1])
	// 选当前节点：取不选孩子的和 + 当前值
	selectNode := left[0] + right[0] + node.Val

	return [2]int{notSelect, selectNode}
}

// ============================================================================
// 工具函数
// ============================================================================

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
