package main

import (
	"container/heap"
	"fmt"
)

// 1091. 二进制矩阵中的最短路径
// 在一个 N × N 的方形网格中，每个单元格有两种状态：空（0）或者阻塞（1）。
// 一条从左上角到右下角、长度为 k 的畅通路径，由满足下述条件的单元格 C_1, C_2, ..., C_k 组成：
//     相邻单元格 C_i 和 C_{i+1} 在八个方向之一上连通（此时，C_i 和 C_{i+1} 不同且共享边或角）
//     C_1 位于 (0, 0)（即，值为 grid[0][0]）
//     C_k 位于 (N-1, N-1)（即，值为 grid[N-1][N-1]）
//     如果 C_i 位于 (r, c)，则 grid[r][c] 为空（即，grid[r][c] == 0）
// 返回这条从左上角到右下角的最短畅通路径的长度。如果不存在这样的路径，返回 -1 。
// 提示：
//     1 <= grid.length == grid[0].length <= 100
//     grid[i][j] 为 0 或 1
// https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/
func main() {
	// 4
	fmt.Println(shortestPathBinaryMatrix2([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}))
	// 6
	fmt.Println(shortestPathBinaryMatrix2([][]int{{0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}, {0, 0, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 1, 0, 0}}))
	//10
	fmt.Println(shortestPathBinaryMatrix2([][]int{{0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1}, {0, 0, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 0, 1, 0, 0, 1, 1, 0}, {0, 0, 0, 0, 0, 1, 0, 0, 0}, {0, 1, 0, 1, 0, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 0, 1, 0}}))
	//  7
	fmt.Println(shortestPathBinaryMatrix2([][]int{{0, 0, 0, 0, 1, 1}, {0, 1, 0, 0, 1, 0}, {1, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1}, {0, 1, 0, 0, 0, 1}, {0, 0, 1, 0, 0, 0}}))
}

// 法一：BFS
var around = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

// 省略了visied数组，直接在grid上进行修改
func shortestPathBinaryMatrix(grid [][]int) (level int) {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	} else if n <= 2 {
		return n
	}
	queue := [][2]int{{0, 0}}
	grid[0][0] = 2
	level++
	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, diff := range around {
				x, y := cur[0]+diff[0], cur[1]+diff[1]
				if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == 0 {
					if x == n-1 && y == n-1 {
						return
					}
					queue = append(queue, [2]int{x, y})
					grid[x][y] = 2
				}
			}

		}
		queue = queue[size:]
	}
	return -1
}

// 法二：A*
// 估价函数h(n)代表从当前点到终点的曼哈顿距离（坐标差绝对值之和）
// 优先级是估价函数的值加上当前点已走的距离，只使用曼哈顿距离只能得到一个较优值
// 使用一个小顶堆取代pq
func shortestPathBinaryMatrix2(grid [][]int) (minDist int) {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	} else if n <= 2 {
		return n
	}
	var pq priorityQueue
	maxPos := n - 1
	pq = append(pq, node{x: 0, y: 0})
	dist := make(map[[2]int]int, n*n)
	dist[[2]int{0, 0}] = 1
	for len(pq) > 0 {
		cur := heap.Pop(&pq).(node)
		if grid[cur.x][cur.y] == 2 {
			continue
		}
		if cur.x == maxPos && cur.y == maxPos {
			return dist[[2]int{maxPos, maxPos}]
		}
		grid[cur.x][cur.y] = 2
		for _, diff := range around {
			x, y := cur.x+diff[0], cur.y+diff[1]
			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] != 1 {
				heap.Push(&pq, node{x, y, heuristicHelper(x, y, maxPos) + dist[[2]int{cur.x, cur.y}] + 1})
				// 剪枝，同一个点有多条到达路径，如果有更短的路径，就更新
				if dist[[2]int{x, y}] == 0 || dist[[2]int{cur.x, cur.y}]+1 < dist[[2]int{x, y}] {
					dist[[2]int{x, y}] = dist[[2]int{cur.x, cur.y}] + 1
				}
			}
		}
	}
	return -1
}

type priorityQueue []node
type node struct {
	x, y     int
	priority int
}

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(node)) }
func (pq *priorityQueue) Pop() interface{} {
	last := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return last
}

// 估价函数取曼哈顿距离
// 这里做了一个优化，直接取最大的边长
// 返回值越低越好
func heuristicHelper(i, j, maxPos int) int {
	return getMax((maxPos - i), (maxPos - j))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
