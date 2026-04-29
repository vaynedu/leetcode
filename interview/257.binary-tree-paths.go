package interview

/*
257. 二叉树的所有路径

给你二叉树 root，返回所有从根到叶子的路径

时间：O(n) | 空间：O(h)
*/

// BinaryTreePaths 返回所有根到叶子路径
func BinaryTreePaths(root *TreeNode) []string {
	var result []string
	var path []int
	pathsRec(root, path, &result)
	return result
}

func pathsRec(node *TreeNode, path []int, result *[]string) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	// 叶子节点：收集路径
	if node.Left == nil && node.Right == nil {
		b := ""
		for i, v := range path {
			if i > 0 {
				b += "->"
			}
			b += itoa(v)
		}
		*result = append(*result, b)
		return
	}
	pathsRec(node.Left, path, result)
	pathsRec(node.Right, path, result)
}

// itoa 简单整数转字符串
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	if neg {
		b = append([]byte{'-'}, b...)
	}
	return string(b)
}
