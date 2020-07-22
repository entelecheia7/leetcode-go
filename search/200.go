package main

import "fmt"

// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
// 岛屿总是被水包围，并且每座岛屿只能由水平方向或竖直方向上相邻的陆地连接形成。
// 此外，你可以假设该网格的四条边均被水包围。
// https://leetcode-cn.com/problems/number-of-islands/
func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	fmt.Println(numIslands2([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
}

// 法一：dfs
func numIslands(grid [][]byte) (count int) {
	if len(grid) == 0 {
		return 0
	}
	lx := len(grid)
	ly := len(grid[0])
	visited := make([][]bool, lx)
	for i := 0; i < lx; i++ {
		visited[i] = make([]bool, ly)
	}
	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				// 第一次发现一个岛的坐标，递归标记周围是‘1’的坐标为已访问
				visited[i][j] = true
				markIsland(grid, i, j, lx, ly, visited)
				count++
			}

		}
	}
	return count
}

var around = [4][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}

func markIsland(grid [][]byte, x, y, lx, ly int, visited [][]bool) {
	if x < 0 || y < 0 || x >= lx || y >= ly {
		return
	}
	for _, diff := range around {
		newX, newY := x+diff[0], y+diff[1]
		if newX >= 0 && newX < lx && newY >= 0 && newY < ly && grid[newX][newY] == '1' && !visited[newX][newY] {
			visited[newX][newY] = true
			markIsland(grid, newX, newY, lx, ly, visited)
		}
	}
}

func numIslands2(grid [][]byte) (count int) {
	if len(grid) == 0 {
		return 0
	}
	lx := len(grid)
	ly := len(grid[0])
	visited := make([][]bool, lx)
	for i := 0; i < lx; i++ {
		visited[i] = make([]bool, ly)
	}

	for i := 0; i < lx; i++ {
		for j := 0; j < ly; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				queue := [][2]int{{i, j}}
				visited[i][j] = true
				for len(queue) > 0 {
					x, y := queue[0][0], queue[0][1]
					for _, diff := range around {
						newX, newY := x+diff[0], y+diff[1]
						if newX >= 0 && newX < lx && newY >= 0 && newY < ly && grid[newX][newY] == '1' && !visited[newX][newY] {
							visited[newX][newY] = true
							queue = append(queue, [2]int{newX, newY})
						}
					}
					queue = queue[1:]
				}

				count++
			}
		}
	}
	return count
}

// 法三：并查集
// 「并查集」主要用于解决「动态连通性」问题，重点关注的是连接问题，不关注路径问题。
// 对于本题，就是将水域和周边水域连接，岛屿和周边岛屿连接
// 岛屿的数量就是岛屿联通集合的数目
// tbc
