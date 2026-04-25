package interview

// 373. 查找和最小的K对数
// 最小堆：pair(i,j) 的和 = nums1[i] + nums2[j]
// 从 (0,0) 开始，把 (i+1,j) 和 (i,j+1) 加入堆（已排序剪枝）
import "container/heap"

type pair struct {
    i, j int
    sum  int
}

type minPairHeap []pair

func (h minPairHeap) Len() int            { return len(h) }
func (h minPairHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h minPairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minPairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *minPairHeap) Pop() interface{} {
    n := len(*h)
    top := (*h)[n-1]
    *h = (*h)[:n-1]
    return top
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
        return nil
    }

    h := &minPairHeap{}
    heap.Init(h)
    visited := make(map[[2]int]bool)

    heap.Push(h, pair{0, 0, nums1[0] + nums2[0]})
    visited[[2]int{0, 0}] = true

    res := [][]int{}
    for h.Len() > 0 && len(res) < k {
        p := heap.Pop(h).(pair)
        res = append(res, []int{nums1[p.i], nums2[p.j]})

        if p.i+1 < len(nums1) && !visited[[2]int{p.i + 1, p.j}] {
            heap.Push(h, pair{p.i + 1, p.j, nums1[p.i+1] + nums2[p.j]})
            visited[[2]int{p.i + 1, p.j}] = true
        }
        if p.j+1 < len(nums2) && !visited[[2]int{p.i, p.j + 1}] {
            heap.Push(h, pair{p.i, p.j + 1, nums1[p.i] + nums2[p.j+1]})
            visited[[2]int{p.i, p.j + 1}] = true
        }
    }
    return res
}
