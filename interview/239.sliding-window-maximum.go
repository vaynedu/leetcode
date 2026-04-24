package interview

// 239. 滑动窗口最大值（Sliding Window Maximum）
// 单调 deque：存下标保持递减，队首是当前窗口最大值的下标

func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0 || k == 0 {
        return []int{}
    }

    var deque []int // 存下标，保持单调递减
    result := []int{}

    for R := 0; R < len(nums); R++ {
        // 1. 新元素入队前：把比它小的从队尾踢掉
        for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[R] {
            deque = deque[:len(deque)-1]
        }
        // 2. 新元素下标加入队尾
        deque = append(deque, R)

        // 3. 踢掉已经滑出窗口的队首（deque[0] 在窗口左边界之外）
        if len(deque) > 0 && deque[0] <= R-k {
            deque = deque[1:]
        }

        // 4. 窗口形成后（至少 k 个元素），记录最大值
        if R >= k-1 {
            result = append(result, nums[deque[0]])
        }
    }

    return result
}
