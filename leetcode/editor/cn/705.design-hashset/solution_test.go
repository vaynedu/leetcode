package leetcode

// 设计哈希集合

import (
	"fmt"
	"testing"
)

//不使用任何内建的哈希表库设计一个哈希集合（HashSet）。
//
// 实现 MyHashSet 类：
//
//
// void add(key) 向哈希集合中插入值 key 。
// bool contains(key) 返回哈希集合中是否存在这个值 key 。
// void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
//
//
// 示例：
//
//
//输入：
//["MyHashSet", "add", "add", "contains", "contains", "add", "contains",
//"remove", "contains"]
//[[], [1], [2], [1], [3], [2], [2], [2], [2]]
//输出：
//[null, null, null, true, false, null, true, null, false]
//
//解释：
//MyHashSet myHashSet = new MyHashSet();
//myHashSet.add(1);      // set = [1]
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(1); // 返回 True
//myHashSet.contains(3); // 返回 False ，（未找到）
//myHashSet.add(2);      // set = [1, 2]
//myHashSet.contains(2); // 返回 True
//myHashSet.remove(2);   // set = [1]
//myHashSet.contains(2); // 返回 False ，（已移除）
//
//
//
// 提示：
//
//
// 0 <= key <= 10⁶
// 最多调用 10⁴ 次 add、remove 和 contains
//
//
// Related Topics 设计 数组 哈希表 链表 哈希函数 👍 362 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
type MyHashSet struct {
	Data []bool // 使用布尔数组存储键值的存在状态
}

// 构造函数
func Constructor() MyHashSet {
	return MyHashSet{
		Data: make([]bool, 1e6+1), // 初始化布尔数组，大小为 10^6 + 1
	}
}

// Add 方法：向哈希集合中插入值 key
func (this *MyHashSet) Add(key int) {
	if key >= 0 && key <= 1e6 { // 确保 key 在有效范围内
		this.Data[key] = true
	}
}

// Remove 方法：将给定值 key 从哈希集合中删除
func (this *MyHashSet) Remove(key int) {
	if key >= 0 && key <= 1e6 { // 确保 key 在有效范围内
		this.Data[key] = false
	}
}

// Contains 方法：返回哈希集合中是否存在这个值 key
func (this *MyHashSet) Contains(key int) bool {
	if key >= 0 && key <= 1e6 { // 确保 key 在有效范围内
		return this.Data[key]
	}
	return false
}
/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
//leetcode submit region end(Prohibit modification and deletion)

func TestDesignHashset(t *testing.T) {
	fmt.Println("come on baby !!!")

	myHashSet := Constructor()
	myHashSet.Add(1)
	myHashSet.Add(2)
	fmt.Println(myHashSet.Contains(1))
	fmt.Println(myHashSet.Contains(3))
	myHashSet.Remove(2)
	fmt.Println(myHashSet.Contains(2))
}
