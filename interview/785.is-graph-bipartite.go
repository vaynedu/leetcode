package interview

// 785. 判断二分图
// BFS 染色：相邻节点染不同颜色，若冲突则非二分图
func isBipartite(graph [][]int) bool {
    n := len(graph)
    color := make([]int, n) // 0=未访问, 1=红, -1=蓝
    for i := 0; i < n; i++ {
        if color[i] != 0 {
            continue
        }
        color[i] = 1
        queue := []int{i}
        for len(queue) > 0 {
            node := queue[0]
            queue = queue[1:]
            for _, neighbor := range graph[node] {
                if color[neighbor] == 0 {
                    color[neighbor] = -color[node]
                    queue = append(queue, neighbor)
                } else if color[neighbor] == color[node] {
                    return false
                }
            }
        }
    }
    return true
}
