package leetcode

// 207.课程表

import (
	"testing"
)

/**
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。

 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如
果要学习课程 ai 则 必须 先学习课程 bi 。


 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。


 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。



 示例 1：


输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

 示例 2：


输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。



 提示：


 1 <= numCourses <= 2000
 0 <= prerequisites.length <= 5000
 prerequisites[i].length == 2
 0 <= ai, bi < numCourses
 prerequisites[i] 中的所有课程对 互不相同


 👍 2218 👎 0

*/

// leetcode submit region begin(Prohibit modification and deletion)
// leetcode submit region begin(Prohibit modification and deletion)
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 建立邻接表和入度数组
	graph := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	// 构建图
	for _, prereq := range prerequisites {
		course, preCourse := prereq[0], prereq[1]
		graph[preCourse] = append(graph[preCourse], course)
		inDegree[course]++
	}

	// 将所有入度为0的课程加入队列
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// 记录可以完成的课程数量
	completed := 0

	// BFS拓扑排序
	for len(queue) > 0 {
		// 取出一门可以学习的课程
		current := queue[0]
		queue = queue[1:]
		completed++

		// 遍历当前课程的后续课程
		for _, nextCourse := range graph[current] {
			// 将后续课程的入度减1
			inDegree[nextCourse]--
			// 如果入度为0，说明可以学习了
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// 如果完成的课程数量等于总课程数，说明可以完成所有课程
	return completed == numCourses
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCourseSchedule(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			name:          "示例1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      true,
		},
		{
			name:          "示例2",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
		{
			name:          "无先修课程",
			numCourses:    3,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "线性依赖",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			expected:      true,
		},
		{
			name:          "复杂依赖无环",
			numCourses:    6,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {5, 3}, {4, 5}},
			expected:      true,
		},
		{
			name:          "复杂依赖有环",
			numCourses:    6,
			prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {5, 3}, {4, 5}, {0, 4}},
			expected:      false,
		},
		{
			name:          "自环",
			numCourses:    1,
			prerequisites: [][]int{{0, 0}},
			expected:      false,
		},
		{
			name:          "单个课程",
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			name:          "多个独立环",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {0, 1}, {3, 2}, {2, 3}},
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建prerequisites的副本，避免修改原始数据
			prereqCopy := make([][]int, len(tt.prerequisites))
			for i := range tt.prerequisites {
				prereqCopy[i] = make([]int, len(tt.prerequisites[i]))
				copy(prereqCopy[i], tt.prerequisites[i])
			}

			result := canFinish(tt.numCourses, prereqCopy)
			if result != tt.expected {
				t.Errorf("canFinish(%d, %v) = %v, expected %v",
					tt.numCourses, tt.prerequisites, result, tt.expected)
			}
		})
	}
}
