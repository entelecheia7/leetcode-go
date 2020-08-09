package main

import "fmt"

// 36. 有效的数独
// 判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。
//     数字 1-9 在每一行只能出现一次。
//     数字 1-9 在每一列只能出现一次。
//     数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
// 	数独部分空格内已填入了数字，空白格用 '.' 表示。
// 	说明:
//     一个有效的数独（部分已被填充）不一定是可解的。
//     只需要根据以上规则，验证已经填入的数字是否有效即可。
//     给定数独序列只包含数字 1-9 和字符 '.' 。
// 	给定数独永远是 9x9 形式的。
// https://leetcode-cn.com/problems/valid-sudoku
func main() {
	fmt.Println(isValidSudoku2([][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}})) // false
}

// 法一：暴力
// 这里的关键是使用 i/3*3+j/3来判断同一个九宫格
func isValidSudoku(board [][]byte) bool {
	var square, row, column [9][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
				squ := i/3*3 + j/3
				if row[i][num] || column[j][num] || square[squ][num] {
					return false
				}
				row[i][num] = true
				column[j][num] = true
				square[squ][num] = true
			}
		}
	}
	return true
}

// 法二：位运算
// 使用一个9位的二进制数来判断一个数字是否被使用过，0为未使用，1为已使用
// best
func isValidSudoku2(board [][]byte) bool {
	var row, column, squ [9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				squNo := i/3*3 + j/3
				num := int(board[i][j] - '1')
				if isUsedNum(num, row[i]) || isUsedNum(num, squ[squNo]) || isUsedNum(num, column[j]) {
					return false
				}
				row[i] = row[i] ^ (1 << num)
				column[j] = column[j] ^ (1 << num)
				squ[squNo] = squ[squNo] ^ (1 << num)
			}
		}
	}
	return true
}

// set表示9位二进制数，n表示需要判断使用的数字
func isUsedNum(n, set int) bool {
	if ((set >> n) & 1) == 1 {
		return true
	}
	return false
}
