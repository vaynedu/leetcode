# 树基础 · 题目索引

> 学完 `tree.go`（前/中/后序递归 + 迭代模板）后，用这些题检验。

## 模板要点

```
前序：根 左 右  — 先处理根，再处理子树
中序：左 根 右  — BST 有序
后序：左 右 根  — 子树处理完再处理根

递归要点：
  - 终止条件：root == nil
  - 递归左右子树
  - （后序）在左右子树之后处理根
```

## 对应题目

| # | 题目 | 难度 | 频率 | 分类 | 题解 |
|---|------|------|------|------|------|
| 144 | [二叉树前序遍历](./../../interview/144.binary-tree-preorder-traversal.go) | 中等 | ★★☆ | 前序 | `go test -run TestPreorder` |
| 94  | [二叉树中序遍历](./../../interview/94.binary-tree-inorder-traversal.go) | 中等 | ★★★ | 中序 | `go test -run TestInorder` |
| 145 | [二叉树后序遍历](./../../interview/145.binary-tree-postorder-traversal.go) | 中等 | ★★☆ | 后序 | `go test -run TestPostorder` |

> 更多二叉树综合题见 [binary-tree/PROBLEMS.md](../binary-tree/PROBLEMS.md)

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/tree/ -v
```
