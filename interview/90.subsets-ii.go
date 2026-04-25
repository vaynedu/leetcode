package interview

// 90. 子集 II
// 回溯 + 排序去重：同层相同值只选第一个
import "sort"

func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    path := []int{}

    var backtrack func(start int)
    backtrack = func(start int) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)

        for i := start; i < len(nums); i++ {
            // 同层剪枝：相同值只选第一个
            if i > start && nums[i] == nums[i-1] {
                continue
            }
            path = append(path, nums[i])
            backtrack(i + 1)
            path = path[:len(path)-1]
        }
    }

    backtrack(0)
    return res
}
