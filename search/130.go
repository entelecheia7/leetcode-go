package main

import "fmt"

// 130. 被围绕的区域
// 给定一个二维的矩阵，包含 'X' 和 'O'（字母 O）。
// 找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
// 解释:
// 被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
// https://leetcode-cn.com/problems/surrounded-regions
func main() {
	board := [][]byte{
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'}}
	solve(board)
	fmt.Println(board)
}

// 法一：dfs，类似200.岛屿数量问题
// 遍历board，如果在边界上遇到O，递归标记连通的点O为V
func solve(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	m, n := len(board), len(board[0])
	// 标记边界的O
	rows := []int{0, m - 1}
	colomns := []int{0, n - 1}
	for _, i := range rows {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				markBorder(board, i, j, m, n)
			}
		}
	}
	for _, j := range colomns {
		for i := 1; i < m-1; i++ {
			if board[i][j] == 'O' {
				markBorder(board, i, j, m, n)
			}
		}
	}
	// 遍历处理所有的O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case 'V':
				board[i][j] = 'O'
			}
		}
	}
}
func markBorder(board [][]byte, i, j, m, n int) {
	if i < 0 || i == m || j < 0 || j == n || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'V'
	markBorder(board, i+1, j, m, n)
	markBorder(board, i-1, j, m, n)
	markBorder(board, i, j+1, m, n)
	markBorder(board, i, j-1, m, n)
}

// 法二：并查集
// 将边界的 O 的集合进行统计，然后遍历board，将不属于边界集合的 O 更改为 X
var around = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func solve2(board [][]byte) {
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	m, n := len(board), len(board[0])
	border := m * n
	parent := make([]int, border+1)
	for i := 0; i <= border; i++ {
		parent[i] = i
	}
	rows := []int{0, m - 1}
	colomns := []int{0, n - 1}
	for _, row := range rows {
		for j := 0; j < n; j++ {
			union(board, row, j, m, n, parent, border)
		}
	}
	for _, colomn := range colomns {
		for i := 1; i < m-1; i++ {
			union(board, i, colomn, m, n, parent, border)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' && parent[i*n+j] != border {
				board[i][j] = 'X'
			}
		}
	}
}
func union(board [][]byte, i, j, m, n int, parent []int, border int) {
	if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' || parent[i*n+j] == border {
		return
	}
	parent[i*n+j] = border
	for _, diff := range around {
		union(board, i+diff[0], j+diff[1], m, n, parent, border)
	}
}
