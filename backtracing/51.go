package main

import (
	"fmt"
	"strings"
)

// 51. N皇后
// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。
// 每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
// 提示：
//     皇后，是国际象棋中的棋子，意味着国王的妻子。皇后只做一件事，那就是“吃子”。当她遇见可以吃的棋子时，就迅速冲上去吃掉棋子。当然，她横、竖、斜都可走一到七步，可进可退。（引用自 百度百科 - 皇后 ）
// https://leetcode-cn.com/problems/n-queens/
func main() {
	fmt.Println(solveNQueens(4))
}

// 回溯
func solveNQueens(n int) (result [][]string) {
	// 生成空棋盘
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	nQueensHelper(n, board, 0, &result)
	return result
}

// row代表放置的行，范围：0~n-1
func nQueensHelper(n int, board []string, row int, result *[][]string) {
	if row == n {
		tmp := make([]string, n)
		copy(tmp, board)
		*result = append(*result, tmp)
		return
	}
	for column := 0; column < n; column++ {
		if checkNQueen(n, row, column, board) {
			origRow := board[row]
			rowStr := []byte(board[row])
			rowStr[column] = 'Q'
			board[row] = string(rowStr)
			nQueensHelper(n, board, row+1, result)
			board[row] = origRow
		}
	}
}

// row和column代表新皇后想要放置的行和列
func checkNQueen(n, row, column int, cur []string) bool {
	leftup, rightup := column-1, column+1
	for i := row - 1; i >= 0; i-- { // i表示已放置皇后的行
		if cur[i][column] == 'Q' { // 检查竖行
			return false
		}
		if leftup >= 0 && cur[i][leftup] == 'Q' { // 检查左斜线
			return false
		}
		if rightup < n && cur[i][rightup] == 'Q' { // 检查右斜线
			return false
		}
		leftup--
		rightup++
	}
	return true
}

// 示例:
// 输入: 4
// 输出: [
//  [".Q..",  // 解法 1
//   "...Q",
//   "Q...",
//   "..Q."],

//  ["..Q.",  // 解法 2
//   "Q...",
//   "...Q",
//   ".Q.."]
// ]
// 解释: 4 皇后问题存在两个不同的解法。
