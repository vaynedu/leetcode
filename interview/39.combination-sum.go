package interview

// 39. 组合总和
// 回溯 + 剪枝：从 candidates 选，可重复选，从 start 开始
func combinationSum(candidates []int, target int) [][]int {
    res := [][]int{}
    path := []int{}

    var backtrack func(start, sum int)
    backtrack = func(start, sum int) {
        if sum == target {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        if sum > target {
            return
        }
        for i := start; i < len(candidates); i++ {
            path = append(path, candidates[i])
            backtrack(i, sum+candidates[i]) // i 而不是 i+1：可重复选当前数
            path = path[:len(path)-1]
        }
    }

    backtrack(0, 0)
    return res
}
