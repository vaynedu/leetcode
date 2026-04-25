package interview

// 295. 数据流中位数
// 两个堆：大顶堆存较小的一半，小顶堆存较大的一半
// 保持 len(maxH) >= len(minH)，中位数 = maxH[0]（奇）或 (maxH[0]+minH[0])/2（偶）
import "container/heap"

type medianFinder struct {
    maxH *maxHeap // 大顶堆：较小的一半
    minH *minHeap // 小顶堆：较大的一半
}

func Constructor295() medianFinder {
    return medianFinder{
        maxH: &maxHeap{},
        minH: &minHeap{},
    }
}

func (m *medianFinder) AddNum(num int) {
    heap.Push(m.maxH, num)
    heap.Push(m.minH, heap.Pop(m.maxH))
    if m.maxH.Len() < m.minH.Len() {
        heap.Push(m.maxH, heap.Pop(m.minH))
    }
}

func (m *medianFinder) FindMedian() float64 {
    if m.maxH.Len() > m.minH.Len() {
        return float64((*m.maxH)[0])
    }
    return float64((*m.maxH)[0]+(*m.minH)[0]) / 2.0
}

// ============== 大顶堆 ==============
type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
    n := len(*h)
    top := (*h)[n-1]
    *h = (*h)[:n-1]
    return top
}

// ============== 小顶堆 ==============
type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
    n := len(*h)
    top := (*h)[n-1]
    *h = (*h)[:n-1]
    return top
}
