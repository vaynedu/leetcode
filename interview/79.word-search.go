package interview

// 79. 单词搜索
// DFS 回溯：从每个格子出发，尝试上下左右匹配
func exist(board [][]byte, word string) bool {
    if len(board) == 0 || len(word) == 0 {
        return false
    }
    m, n := len(board), len(board[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    var dfs func(i, j, idx int) bool
    dfs = func(i, j, idx int) bool {
        if idx == len(word) {
            return true
        }
        if i < 0 || i >= m || j < 0 || j >= n {
            return false
        }
        if visited[i][j] || board[i][j] != word[idx] {
            return false
        }
        visited[i][j] = true
        defer func() { visited[i][j] = false }()
        return dfs(i+1, j, idx+1) || dfs(i-1, j, idx+1) ||
               dfs(i, j+1, idx+1) || dfs(i, j-1, idx+1)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}
