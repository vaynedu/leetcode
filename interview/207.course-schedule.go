package interview

// 207. 课程表
// BFS 拓扑排序：统计入度，有环则无法完成
func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    inDegree := make([]int, numCourses)
    for _, p := range prerequisites {
        a, b := p[0], p[1] // a 依赖 b
        graph[b] = append(graph[b], a)
        inDegree[a]++
    }

    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    visited := 0
    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        visited++
        for _, next := range graph[course] {
            inDegree[next]--
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return visited == numCourses
}
