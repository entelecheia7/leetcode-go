package main

import "fmt"

// 37. 解数独
// 编写一个程序，通过已填充的空格来解决数独问题。
// 一个数独的解法需遵循如下规则：
//     数字 1-9 在每一行只能出现一次。
//     数字 1-9 在每一列只能出现一次。
//     数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
// 	Note:
//     给定的数独序列只包含数字 1-9 和字符 '.' 。
//     你可以假设给定的数独只有唯一解。
// 	给定数独永远是 9x9 形式的。
// https://leetcode-cn.com/problems/sudoku-solver
func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	fmt.Println(board)
}

// 法一：回溯+位运算
func solveSudoku(board [][]byte) {
	var row, line, squ [9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				row[i] |= 1 << num
				line[j] |= 1 << num
				squNo := i/3*3 + j/3
				squ[squNo] |= 1 << num
			}
		}
	}
	sudokuHelper(board, 0, 0, row, line, squ)
}
func sudokuHelper(board [][]byte, i, j int, row, line, squ [9]int) bool {
	if j == 9 {
		return sudokuHelper(board, i+1, 0, row, line, squ)
	}
	if i == 9 {
		return true
	}
	if board[i][j] != '.' {
		return sudokuHelper(board, i, j+1, row, line, squ)
	}
	squNo := i/3*3 + j/3
	for b := '1'; b <= '9'; b++ {
		n := int(b - '1')
		if !checkUsedNum(n, row[i], line[j], squ[squNo]) {
			offset := 1 << n
			board[i][j] = byte(b)
			row[i] |= offset
			line[j] |= offset
			squ[squNo] |= offset
			if sudokuHelper(board, i, j+1, row, line, squ) {
				return true
			}
			board[i][j] = '.'
			row[i] ^= offset
			line[j] ^= offset
			squ[squNo] ^= offset
		}
	}
	return false

}
func checkUsedNum(n, row, col, squ int) bool {
	if ((row>>n)&1) == 1 || ((col>>n)&1) == 1 || ((squ>>n)&1) == 1 {
		return true
	}
	return false
}
