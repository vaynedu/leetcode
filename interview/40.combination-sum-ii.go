package interview

// 40. 组合总和 II
// 回溯 + 排序去重：每个数只能选一次，同层不能选相同的数
import "sort"

func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    res := [][]int{}
    path := []int{}

    var backtrack func(pos, sum int)
    backtrack = func(pos, sum int) {
        if sum == target {
            tmp := make([]int, len(path))
            copy(tmp, path)
            res = append(res, tmp)
            return
        }
        for i := pos; i < len(candidates); i++ {
            // 同层剪枝：相同值只选第一个
            if i > pos && candidates[i] == candidates[i-1] {
                continue
            }
            if sum+candidates[i] > target {
                break // 排序后剪枝
            }
            path = append(path, candidates[i])
            backtrack(i+1, sum+candidates[i])
            path = path[:len(path)-1]
        }
    }

    backtrack(0, 0)
    return res
}
