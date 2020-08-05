package main

import (
	"fmt"
)

// 207. 课程表
// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
// 提示：
//     输入的先决条件是由 边缘列表 表示的图形，而不是 邻接矩阵 。详情请参见图的表示法。
//     你可以假定输入的先决条件中没有重复的边。
//     1 <= numCourses <= 10^5
// https://leetcode-cn.com/problems/course-schedule
func main() {
	// fmt.Println(canFinish2(2, [][]int{{1, 0}}))         // true
	// fmt.Println(canFinish2(2, [][]int{{0, 1}, {1, 0}})) // false
	// fmt.Println(canFinish2(3, [][]int{{1, 0}, {2, 1}}))         // true
	fmt.Println(canFinish2(3, [][]int{{0, 1}, {0, 2}, {1, 2}})) // true
}

// 课程依赖相当于一个有向无环图，如果环，说明失败

// 法一：拓扑排序-Kahn 算法（bfs）
// 统计顶点的入度，如果某个顶点入度（表示有多少条边指向这个顶点）为 0， 也就表示，没有任何顶点必须先于这个顶点执行，那么这个顶点就可以执行了。
// 将入度为0的顶点删除，同时将它指向的边的入度-1
// 循环执行以上过程，看是否可以完全输出
func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		indegree[edge[1]]++
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}
	queue := []int{} // 统计入度为0的顶点
	for k, v := range indegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	sorted := 0
	for len(queue) > 0 {
		next := graph[queue[0]]
		queue = queue[1:]
		sorted++
		for _, v := range next {
			indegree[v]--
			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	return sorted == numCourses
}

// 法二：拓扑排序-dfs
// 1. 构建邻接表
// 2. dfs访问
// 将每个点的状态分为三类：0未访问，1 代表搜索中，该顶点的出度尚未搜索完毕，2表示该顶点及其出度都搜索完毕
// 如果在搜索中，再次遇到状态为 1 的顶点，说明有环
func canFinish2(numCourses int, prerequisites [][]int) (result bool) {
	graph := make([][]int, numCourses)
	for _, edge := range prerequisites {
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	result = true
	visited := make([]int8, numCourses)
	for i := 0; i < numCourses; i++ {
		if visited[i] == 0 {
			dfsHelper(i, graph, visited, &result)
		}
	}
	return result
}
func dfsHelper(cur int, graph [][]int, visited []int8, result *bool) {
	if !(*result) {
		return
	}
	visited[cur] = 1
	next := graph[cur]
	for _, v := range next {
		if visited[v] == 1 {
			*result = false
			return
		}
		if visited[v] != 2 {
			dfsHelper(v, graph, visited, result)
		}
	}
	visited[cur] = 2
}
