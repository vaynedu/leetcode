# 二分查找边界详解：区间与循环条件

## 识别信号

- 输入本身有序，或者局部有序。
- 题目要求找某个位置、边界、第一个满足条件、最小可行值。
- 答案具备单调性：某个条件从 false 变 true，或从 true 变 false。

## 核心不变量

- 搜索区间里始终保留可能答案。
- 每次比较后排除的一半一定不含答案。
- 使用左闭右开 `[left,right)` 时，退出条件 `left == right`，答案通常就是 `left`。
- 使用闭区间 `[left,right]` 时，退出条件 `left > right`，需要根据题意返回 `left`、`right` 或命中位置。

## 决策流程

```text
1. 先判断是找值、找左边界、找右边界，还是在答案空间上二分。
2. 明确区间定义：[l,r] 或 [l,r)。
3. 写出 check(mid) 的含义。
4. 决定 mid 满足条件时保留哪一边。
5. 退出后用不变量解释返回值。
```

## 为什么不会漏答案

二分不是“猜中点”，而是每次用单调性证明一半区间不可能含答案。只要收缩时保留了可能答案，区间缩到空或单点时就得到边界。

## 混乱根源

把两件事混在一起了：

```
区间定义  [left, right]  或  [left, right)
循环条件  left <= right  或  left < right
```

**区间定义**决定 `right` 一开始等于什么。
**循环条件**决定循环什么时候停。
**收缩方式**决定 `right = mid - 1` 还是 `right = mid`。

---

## 一、区间定义的两种方式

### 闭区间 [left, right]

```
[left, right] 表示 left 和 right 都是有效位置

例：[0, 5] → 有效位置是 0, 1, 2, 3, 4, 5（6个位置）
```

### 左闭右开 [left, right)

```
[left, right) 表示 left 有效，right 无效（right 是"左边界外的第一个位置"）

例：[0, 5) → 有效位置是 0, 1, 2, 3, 4（5个位置）
```

**为什么要右开？** 因为这样 right 可以直接等于 `len(nums)`，不需要 `-1`，边界更干净。

---

## 二、两种区间的循环条件

### [left, right] 闭区间 → 用 `left <= right`

```
while (left <= right)     // left == right 时还有1个位置要检查
```

退出时 `left > right`（区间空了）

### [left, right) 左闭右开 → 用 `left < right`

```
while (left < right)     // left == right 时区间为空，退出
```

退出时 `left == right`（区间空了）

---

## 三、四种情况对比

```
┌──────────────┬────────────┬─────────────────┬──────────────────┐
│ 类型         │ 区间       │ 循环条件         │ right 变化        │
├──────────────┼────────────┼─────────────────┼──────────────────┤
│ 标准二分     │ [l, r]     │ left <= right    │ right = mid - 1  │
│ 找target    │ 闭区间     │                  │ left  = mid + 1  │
├──────────────┼────────────┼─────────────────┼──────────────────┤
│ 左边界       │ [l, r)     │ left < right     │ right = mid      │
│ 第一个>=target│ 左闭右开  │                  │ left  = mid + 1  │
├──────────────┼────────────┼─────────────────┼──────────────────┤
│ 右边界       │ [l, r)     │ left < right     │ right = mid      │
│ 第一个>target │ 左闭右开  │                  │ left  = mid + 1  │
└──────────────┴────────────┴─────────────────┴──────────────────┘
```

---

## 四、左边界（第一个 >= target）详解

**目标**：找第一个 `>= target` 的位置。

### 区间：[left, right)，right = len(nums)

```
初始：left = 0, right = len(nums)
例：[5,7,7,8,8,10], len=6
初始：left=0, right=6
```

### 循环：`left < right`

```
退出时 left == right == 第一个>=target的位置
```

### 收缩规则：

```
         ┌── nums[mid] < target
         │   → target 在 mid 的右边
         │   → left = mid + 1（mid本身不是答案）
         │
nums[mid] 的情况
         │
         └── nums[mid] >= target
             → target 在 mid 的左边（含mid本身）
             → right = mid（mid可能是答案，保留mid位置）
```

### 关键：为什么 right = mid 而不是 mid - 1？

因为是**左闭右开**，`mid` 位置本身可能是答案。
- `right = mid - 1` 会**跳过正确答案**
- `right = mid` 保留答案的可能性

---

## 五、右边界（第一个 > target）详解

**目标**：找第一个 `> target` 的位置。

### 收缩规则（左闭右开）：

```
         ┌── nums[mid] <= target
         │   → target 在 mid 的左边（mid本身<=target，不是答案）
         │   → left = mid + 1（跳过mid）
         │
nums[mid] 的情况
         │
         └── nums[mid] > target
             → target 在 mid 的左边，mid 可能是答案
             → right = mid
```

### 与左边界的唯一区别：比较符号

```
左边界：nums[mid] < target  → left = mid + 1
        else               → right = mid

右边界：nums[mid] <= target → left = mid + 1
        else               → right = mid
```

---

## 六、左边界找 target 的例子

**数组 [5,7,7,8,8,10]，target = 8**

```
初始：left=0, right=6, [left,right) = [0,6)

第1轮：mid=3, nums[3]=8
  nums[3]=8 >= 8 → right = 3
  [0,3) = [0,2] → left=0, right=3

第2轮：mid=1, nums[1]=7
  nums[1]=7 < 8  → left = 2
  [2,3) = [2,2] → left=2, right=3

第3轮：mid=2, nums[2]=7
  nums[2]=7 < 8  → left = 3
  [3,3) → left == right，退出

left = 3 ← 第一个 >= 8 的位置 ✓
```

---

## 七、右边界找 target 的例子

**数组 [5,7,7,8,8,10]，target = 8**

找第一个 `> 8` 的位置（用 lowerBound(9)）：
```
初始：left=0, right=6

第1轮：mid=3, nums[3]=8
  nums[3]=8 <= 9 → left = 4
  [4,6)

第2轮：mid=5, nums[5]=10
  nums[5]=10 > 9 → right = 5
  [4,5)

第3轮：mid=4, nums[4]=8
  nums[4]=8 <= 9 → left = 5
  [5,5) → 退出

left = 5 = lowerBound(9)
right = 5 - 1 = 4 ← 最后一个 <= 8 的位置 ✓
```

---

## 八、什么时候开，什么时候闭？

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  什么时候闭区间 [left, right]？                              │
│    → 标准二分，找存在的 target                               │
│    → left = mid + 1, right = mid - 1（两个都收缩）          │
│                                                             │
│  什么时候开区间 [left, right)？                              │
│    → 找边界（第一个 >= 或 第一个 >）                        │
│    → left = mid + 1, right = mid（只有 left 收缩 +1）       │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

---

## 九、一句话总结

```
闭区间 [l,r] 标准二分：  while(left <= right)  right=mid-1 left=mid+1
左闭右开 [l,r) 边界查找：while(left < right)   right=mid    left=mid+1

左边界：找 >= target，条件是 nums[mid] < target → left=mid+1
右边界：找 > target，条件是 nums[mid] <= target → left=mid+1
        （唯一区别：= 时往左还是往右）
```

---

## 十、为什么用左闭右开？

因为找边界时，`mid` 本身可能是答案，不能用 `mid - 1` 跳过它。

闭区间 `[l,r]` 用 `right = mid - 1` 是因为：
- 此时 `nums[mid] != target` 已经确定
- `mid` 不是答案，所以可以 -1

左闭右开 `[l,r)` 用 `right = mid` 是因为：
- 此时还不确定 `mid` 是不是答案
- 保留 `mid` 位置，让它还有机会被选中

---

## 十一、一句话总结（面试速背）

```
找「存在的值」→ 闭区间，代码更简单
找「边界位置」→ 左闭右开，边界处理更干净

实战中记住：

for left <= right        // 闭区间，找到了就 return
for left < right         // 左闭右开，收缩到 left == right

right = mid - 1          // 闭区间，mid已验证不是答案
right = mid              // 左闭右开，mid可能是答案

面试被问到选哪个：

标准回答：
"都可以实现，但左闭右开在找边界时更不容易出错，
因为 right 初始值可以直接用 len(nums)，不需要 len(nums)-1。"
```
