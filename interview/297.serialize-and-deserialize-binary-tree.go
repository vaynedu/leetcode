package interview

import (
    "strconv"
    "strings"
)

// 297. 二叉树的序列化与反序列化
// 前序遍历序列化，空节点用 # 分隔
type Codec struct{}

func NewCodec() Codec {
    return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "#"
    }
    left := c.serialize(root.Left)
    right := c.serialize(root.Right)
    return strconv.Itoa(root.Val) + "," + left + "," + right
}

func (c *Codec) deserialize(data string) *TreeNode {
    if data == "" {
        return nil
    }
    nodes := strings.Split(data, ",")
    idx := 0
    var build func() *TreeNode
    build = func() *TreeNode {
        if idx >= len(nodes) {
            return nil
        }
        val := nodes[idx]
        idx++
        if val == "#" {
            return nil
        }
        v, _ := strconv.Atoi(val)
        node := &TreeNode{Val: v}
        node.Left = build()
        node.Right = build()
        return node
    }
    return build()
}
