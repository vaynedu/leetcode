package interview

// 133. 克隆图
// BFS + 哈希表记录已克隆节点
func cloneGraph(node *GraphNode) *GraphNode {
    if node == nil {
        return nil
    }
    visited := map[*GraphNode]*GraphNode{}
    queue := []*GraphNode{node}
    visited[node] = &GraphNode{Val: node.Val}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, neighbor := range cur.Neighbors {
            if visited[neighbor] == nil {
                visited[neighbor] = &GraphNode{Val: neighbor.Val}
                queue = append(queue, neighbor)
            }
            visited[cur].Neighbors = append(visited[cur].Neighbors, visited[neighbor])
        }
    }
    return visited[node]
}

// GraphNode 定义（避免与 146.lru-cache.Node 冲突）
type GraphNode struct {
    Val       int
    Neighbors []*GraphNode
}
