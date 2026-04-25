package interview

// 78. 子集
// 回溯：每层决定每个元素选或不选，树高 n，2^n 个叶子
func subsets(nums []int) [][]int {
    res := [][]int{}
    path := []int{}

    var backtrack func(start int)
    backtrack = func(start int) {
        // 每层都算一个子集（叶子之前的所有节点都是子集）
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)

        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            backtrack(i + 1)
            path = path[:len(path)-1]
        }
    }

    backtrack(0)
    return res
}
