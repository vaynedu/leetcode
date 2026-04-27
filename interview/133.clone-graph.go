package interview

// 133. 克隆图
// BFS + 哈希表记录已克隆节点
func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    visited := map[*Node]*Node{}
    queue := []*Node{node}
    visited[node] = &Node{Val: node.Val}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]
        for _, neighbor := range cur.Neighbors {
            if visited[neighbor] == nil {
                visited[neighbor] = &Node{Val: neighbor.Val}
                queue = append(queue, neighbor)
            }
            visited[cur].Neighbors = append(visited[cur].Neighbors, visited[neighbor])
        }
    }
    return visited[node]
}

// Node 定义
type Node struct {
    Val       int
    Neighbors []*Node
}
