package interview

// SearchMatrix 74. 搜索二维矩阵
// 难度：Medium | 标签：数组 | 二分查找
// 核心思路：将二维矩阵当作一维数组做二分，或两次二分（先找行，再找列）
// 时间：O(log(m*n)) | 空间：O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	// 将二维坐标映射为一维：index = i * n + j
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)/2
		row := mid / n
		col := mid % n
		val := matrix[row][col]
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
