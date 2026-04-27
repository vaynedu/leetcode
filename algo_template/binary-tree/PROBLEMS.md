# 二叉树 · 题目索引

> 学完 `traversal.go` 模板后，用这些题实战检验。

## 模板要点

```
遍历模板：
  前序：根 左 右   — 访问顺序先于子树
  中序：左 根 右   — BST 有序
  后序：左 右 根   — 子树先于根处理

递归模板：
  func dfs(root):
      if root == nil: return
      // 前序位置
      dfs(root.Left)
      // 中序位置
      dfs(root.Right)
      // 后序位置

路径问题：
  112: 判断是否存在路径（boolean）
  113: 收集所有路径（回溯）
  257: 所有根到叶路径（字符串）
  437: 路径和 III（前缀和）
```

## 对应题目

### 遍历类（6道）

| # | 题目 | 难度 | 频率 | 分类 | 题解 |
|---|------|------|------|------|------|
| 102 | [二叉树层序遍历](./../../interview/102.binary-tree-level-order-traversal.go) | 中等 | ★★★ | BFS层序 | `go test -run TestLevelOrder` |
| 94  | [二叉树中序遍历](./../../interview/94.binary-tree-inorder-traversal.go) | 中等 | ★★★ | 中序 | `go test -run TestInorder` |
| 144 | [二叉树前序遍历](./../../interview/144.binary-tree-preorder-traversal.go) | 中等 | ★★☆ | 前序 | `go test -run TestPreorder` |
| 145 | [二叉树后序遍历](./../../interview/145.binary-tree-postorder-traversal.go) | 中等 | ★★☆ | 后序 | `go test -run TestPostorder` |

### 路径类（4道）

| # | 题目 | 难度 | 频率 | 分类 | 题解 |
|---|------|------|------|------|------|
| 112 | [路径总和](./../../interview/112.path-sum.go) | 简单 | ★★★ | boolean | `go test -run TestHasPathSum` |
| 113 | [路径总和 II](./../../interview/113.path-sum-ii.go) | 中等 | ★★★ | 收集所有 | `go test -run TestPathSumII` |
| 257 | [二叉树所有路径](./../../interview/257.binary-tree-paths.go) | 简单 | ★★★ | 根到叶 | `go test -run TestBinaryTreePaths` |
| 437 | [路径总和 III](./../../interview/437.path-sum-iii.go) | 中等 | ★★★ | 前缀和 | `go test -run TestPathSumIII` |

### 属性/操作类（7道）

| # | 题目 | 难度 | 频率 | 分类 | 题解 |
|---|------|------|------|------|------|
| 572 | [另一棵树的子树](./../../interview/572.subtree-of-another-tree.go) | 简单 | ★★☆ | 子树判定 | `go test -run TestIsSubtree` |
| 617 | [合并二叉树](./../../interview/617.merge-two-binary-trees.go) | 简单 | ★★★ | 同步遍历 | `go test -run TestMergeTrees` |
| 236 | [最近公共祖先](./../../interview/236.lowest-common-ancestor.go) | 中等 | ★★★ | LCA后序 | `go test -run TestLowestCommonAncestor` |
| 1644 | [二叉树最近公共祖先 II](./../../interview/1644.lowest-common-ancestor-ii.go) | 中等 | ★★☆ | LCA+parent | `go test -run TestLowestCommonAncestorII` |

### Hard 综合（3道）

| # | 题目 | 难度 | 频率 | 分类 | 题解 |
|---|------|------|------|------|------|

## 学习路径建议

```
第一轮：104/226/101（套模板，3道简单题）
  ↓
第二轮：112/257/617（路径/合并，dfs套用）
  ↓
第三轮：102/94/98/144/145（遍历+BST验证）
  ↓
第四轮：236/437/543（后序三剑客：LCA/前缀和/直径）
  ↓
第五轮：113/124/105/114（hard门槛，综合型）
```

## 模板测试

```bash
cd ~/go/src/leetcode
go test ./algo_template/binary-tree/ -v
```
