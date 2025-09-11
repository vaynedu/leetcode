package leetcode

// 347.前 K 个高频元素

import (
	"container/heap"
	"fmt"
	"testing"
)

/**
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。



 示例 1：


 输入：nums = [1,1,1,2,2,3], k = 2


 输出：[1,2]

 示例 2：


 输入：nums = [1], k = 1


 输出：[1]

 示例 3：


 输入：nums = [1,2,1,2,1,2,3,1,3,2], k = 2


 输出：[1,2]



 提示：


 1 <= nums.length <= 10⁵
 k 的取值范围是 [1, 数组中不相同的元素的个数]
 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的




 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。

 👍 2067 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type Item struct {
	Val  int // 数值
	Freq int // 频次
}

type MinHeap []*Item

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// 1. 统计每个元素的频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}
	// 2. 创建最小堆
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	// 3. 遍历频率表，将元素加入最小堆
	for value, freq := range freqMap {
		if minHeap.Len() < k {
			heap.Push(minHeap, &Item{
				Val:  value,
				Freq: freq,
			})
		} else if (*minHeap)[0].Freq < freq {
			heap.Pop(minHeap)
			heap.Push(minHeap, &Item{
				Val:  value,
				Freq: freq,
			})
		}
	}
	// 4.从堆中提取结果
	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(minHeap).(*Item).Val
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

// 桶排序方法实现
func topKFrequentBucket(nums []int, k int) []int {
	// 1. 统计每个元素的频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. 创建桶，索引表示频率，值表示该频率的元素列表
	buckets := make([][]int, len(nums)+1)
	for value, freq := range freqMap {
		buckets[freq] = append(buckets[freq], value)
	}

	// 3. 从高频率到低频率收集结果
	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if buckets[i] != nil {
			result = append(result, buckets[i]...)
		}
	}

	// 4. 如果结果超过k个，截取前k个
	if len(result) > k {
		result = result[:k]
	}

	return result
}

func TestTopKFrequentElements(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
