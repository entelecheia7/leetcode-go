package main

import (
	"fmt"
)

// 79. 单词搜索
// 给定一个二维网格和一个单词，找出该单词是否存在于网格中。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
// 提示：
//     board 和 word 中只包含大写和小写英文字母。
//     1 <= board.length <= 200
//     1 <= board[i].length <= 200
//     1 <= word.length <= 10^3
// https://leetcode-cn.com/problems/word-search
func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE")) // true
}

// 使用board来代替visited数组，节省空间
func exist(board [][]byte, word string) (result bool) {
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			existHelper(board, word, 0, i, j, &result)
		}
	}
	return result
}
func existHelper(board [][]byte, word string, k, i, j int, result *bool) {
	if *result || i == len(board) || i < 0 || j == len(board[0]) || j < 0 || board[i][j] == '.' || board[i][j] != word[k] {
		return
	}
	tmp := board[i][j]
	board[i][j] = '.'
	k++
	if k == len(word) {
		*result = true
		return
	}
	existHelper(board, word, k, i+1, j, result)
	existHelper(board, word, k, i-1, j, result)
	existHelper(board, word, k, i, j+1, result)
	existHelper(board, word, k, i, j-1, result)
	board[i][j] = tmp
}
