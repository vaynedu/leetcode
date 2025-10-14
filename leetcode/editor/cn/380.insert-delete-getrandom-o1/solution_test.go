package leetcode

// 380.O(1) 时间插入、删除和获取随机元素

import (
	"math/rand"
	"testing"
)

/**
实现RandomizedSet 类：




 RandomizedSet() 初始化 RandomizedSet 对象
 bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
 bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
 int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。




 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。



 示例：


输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert",
 "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。




 提示：


 -2³¹ <= val <= 2³¹ - 1
 最多调用 insert、remove 和 getRandom 函数 2 * 10⁵ 次
 在调用 getRandom 方法时，数据结构中 至少存在一个 元素。


 Related Topics 设计 数组 哈希表 数学 随机化 👍 973 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:    make([]int, 0),
		indices: make(map[int]int),
	}

}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exist := this.indices[val]; exist {
		return false
	}

	this.nums = append(this.nums, val)
	this.indices[val] = len(this.nums) - 1

	return true

}

func (this *RandomizedSet) Remove(val int) bool {

	index, exists := this.indices[val]
	if !exists {
		return false
	}

	// 获取数组最后一个元素
	lastElement := this.nums[len(this.nums)-1]

	// 将要删除的元素替换为最后一个元素
	this.nums[index] = lastElement

	// 更新最后一个元素的索引
	this.indices[lastElement] = index

	// 删除数组最后一个元素
	this.nums = this.nums[:len(this.nums)-1]
	// 从哈希表中删除该元素
	delete(this.indices, val)

	return true

}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.nums))
	return this.nums[randomIndex]

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)

// 测试函数实现
func TestInsertDeleteGetrandomO1(t *testing.T) {
	// 创建RandomizedSet实例
	rs := Constructor()

	// 测试插入
	if !rs.Insert(1) {
		t.Errorf("Insert 1 should return true")
	}

	if rs.Insert(1) {
		t.Errorf("Insert 1 again should return false")
	}

	if !rs.Insert(2) {
		t.Errorf("Insert 2 should return true")
	}

	// 测试删除
	if rs.Remove(3) {
		t.Errorf("Remove 3 should return false")
	}

	if !rs.Remove(1) {
		t.Errorf("Remove 1 should return true")
	}

	// 测试随机获取
	result := rs.GetRandom()
	if result != 2 {
		// 由于只有一个元素2，随机获取应该返回2
		t.Errorf("GetRandom should return 2, got %d", result)
	}

}
