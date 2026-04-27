# 回溯 · 题目索引

> 学完 `backtrack.go` 模板后，用这些题实战检验。

## 模板要点

```
backtrack(路径, 选择列表):
    if 满足条件:   结果.push(路径.copy())
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 新选择列表)
        撤销选择

去重三定律：
  组合/子集：排序 + if i>pos && nums[i]==nums[i-1] → 跳过
  全排列：used[i-1]==false（同层）
  全排列II：排序 + used[i-1]==false（同层）
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 17 | [电话号码的字母组合](./../../interview/17.letter-combinations-of-a-phone-number.go) | 中等 | ★★☆ | 固定长度 | `go test -run TestLetterCombinations` |
| 22 | [括号生成](./../../interview/22.generate-parentheses.go) | 中等 | ★★★ | 合法括号生成 | `go test -run TestGenerateParenthesis` |
| 39 | [组合总和](./../../interview/39.combination-sum.go) | 中等 | ★★★ | 可重复组合 | `go test -run TestCombinationSum` |
| 40 | [组合总和 II](./../../interview/40.combination-sum-ii.go) | 中等 | ★★★ | 同层去重 | `go test -run TestCombinationSum2` |
| 46 | [全排列](./../../interview/46.permutations.go) | 中等 | ★★★ | 全排列 | `go test -run TestPermute` |
| 47 | [全排列 II](./../../interview/47.permutations-ii.go) | 中等 | ★★★ | used+同层去重 | `go test -run TestPermuteUnique` |
| 78 | [子集](./../../interview/78.subsets.go) | 中等 | ★★★ | 子集 2^n | `go test -run TestSubsets` |
| 90 | [子集 II](./../../interview/90.subsets-ii.go) | 中等 | ★★★ | 排序+同层去重 | `go test -run TestSubsetsWithDup` |

## 学习路径建议

```
第一步：17 / 78 / 46（子集/排列/组合基本型）
  ↓
第二步：39（可重复组合，start不变）
  ↓
第三步：22（括号生成，约束剪枝）
  ↓
第四步：40 / 47 / 90（去重三定律实战）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/backtrack/ -v
```
