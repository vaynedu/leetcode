package leetcode

// 295.数据流的中位数

import (
	"fmt"
	"testing"
)

/**
中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。


 例如 arr = [2,3,4] 的中位数是 3 。
 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。


 实现 MedianFinder 类:


 MedianFinder() 初始化 MedianFinder 对象。
 void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
 double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10⁻⁵ 以内的答案将被接受。


 示例 1：


输入
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
输出
[null, null, null, 1.5, null, 2.0]

解释
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0

 提示:


 -10⁵ <= num <= 10⁵
 在调用 findMedian 之前，数据结构中至少有一个元素
 最多 5 * 10⁴ 次调用 addNum 和 findMedian


 👍 1159 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
import "container/heap"

// 定义最大堆
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // 最大堆
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 定义最小堆
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 最小堆
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// leetcode submit region begin(Prohibit modification and deletion)
type MedianFinder struct {
	left  *MaxHeap // 最大堆，存储较小的一半元素
	right *MinHeap // 最小堆，存储较大的一半元素
}

func Constructor() MedianFinder {
	left := &MaxHeap{}
	right := &MinHeap{}
	heap.Init(left)
	heap.Init(right)

	return MedianFinder{
		left:  left,
		right: right,
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 如果左堆为空或者num小于等于左堆最大值，插入左堆
	if this.left.Len() == 0 || num <= (*this.left)[0] {
		heap.Push(this.left, num)
	} else {
		heap.Push(this.right, num)
	}

	// 平衡两个堆的大小
	// 左堆大小不能超过右堆大小+1
	if this.left.Len() > this.right.Len()+1 {
		heap.Push(this.right, heap.Pop(this.left))
	}
	// 右堆大小不能超过左堆大小
	if this.right.Len() > this.left.Len() {
		heap.Push(this.left, heap.Pop(this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	// 如果总数量为奇数，中位数是左堆（较大堆）的堆顶
	if this.left.Len() > this.right.Len() {
		return float64((*this.left)[0])
	}

	// 如果总数量为偶数，中位数是两个堆顶的平均值
	return float64((*this.left)[0]+(*this.right)[0]) / 2.0
}

//leetcode submit region end(Prohibit modification and deletion)

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestFindMedianFromDataStream(t *testing.T) {
	fmt.Println("come on baby !!!")
	// 生成函数测试用例
	// 要求 有多组tests，并且有输入值，预期值，如果实际返回值和预期值不同，打印错误日志
}
