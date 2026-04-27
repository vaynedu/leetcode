# 二分查找 · 题目索引

> 学完 `binary_search.go` 模板后，用这些题实战检验。

## 模板要点

```
标准二分：left≤right，左中位
找左边界：left<right，right=mid
找右边界：left<right，left=mid+1
旋转数组：先判断哪半有序，再缩小区间

四姐妹：33(搜索) / 35(插入) / 69(Sqrt) / 74(二维)
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类标签 | 题解 |
|---|------|------|------|---------|------|
| 33 | [搜索旋转数组](./../../interview/33.search-in-rotated-sorted-array.go) | 中等 | ★★★ | 旋转数组 | `go test -run TestSearchRotated` |
| 34 | [二分查找边界](./../../interview/34.find-first-and-last-position-of-element-in-sorted-array.go) | 困难 | ★★★ | 左右边界 | `go test -run TestBoundaries` |
| 35 | [搜索插入位置](./../../interview/35.search-insert-position.go) | 简单 | ★★★ | 左边界变体 | `go test -run TestSearchInsert` |
| 69 | [x 的平方根](./../../interview/69.sqrt-x.go) | 简单 | ★★☆ | 右边界变体 | `go test -run TestMySqrt` |
| 74 | [搜索二维矩阵](./../../interview/74.search-a-2d-matrix.go) | 中等 | ★★★ | 旋转+二维 | `go test -run TestSearchMatrix` |
| 81 | [搜索旋转数组 II](./../../interview/81.search-in-rotated-sorted-array-ii.go) | 中等 | ★★☆ | 旋转数组+重复 | `go test -run TestSearchRotatedII` |
| 153 | [寻找旋转数组最小值](./../../interview/153.find-minimum-in-rotated-sorted-array.go) | 中等 | ★★★ | 旋转找最小 | `go test -run TestFindMin` |
| 154 | [寻找旋转数组最小值 II](./../../interview/154.find-minimum-in-rotated-sorted-array-ii.go) | 困难 | ★★☆ | 旋转+重复 | `go test -run TestFindMinII` |

## 学习路径建议

```
第一步：35 / 69（标准二分的左右边界变体）
  ↓
第二步：33 / 81 / 153 / 154（旋转数组四姐妹）
  ↓
第三步：74（二维矩阵二分）
  ↓
第四步：34（二分边界，最核心也最容易错）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/binary-search/ -v
```
