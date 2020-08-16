package main

import (
	"fmt"
)

// 52. N皇后 II
// n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给定一个整数 n，返回 n 皇后不同的解决方案的数量。
// https://leetcode-cn.com/problems/n-queens-ii/
func main() {
	fmt.Println(totalNQueens(4)) //2
}

// 法一：位运算
func totalNQueens(n int) (count int) {
	if n == 1 {
		return 1
	}
	if n <= 3 {
		return
	}
	totalNQueensHelper(n, 0, 0, 0, 0, &count)
	return count
}

// 回溯函数
// col, leftDiagonal, rightDiagonal 分别表示在row这一行皇后在垂直、左斜线和右斜线的攻击范围的攻击范围（攻击范围是1）
func totalNQueensHelper(n, row, col, leftDiagonal, rightDiagonal int, count *int) {
	if row == n {
		(*count)++
		return
	}
	// (1 << n) - 1 将n皇后不需要的高位全部赋为0
	available := (^(col | leftDiagonal | rightDiagonal)) & ((1 << n) - 1) // 当前行的可用位置
	for available != 0 {
		pos := available & -available           // 获取最低位的1的位置
		available = available & (available - 1) // 将pos位置置为0，也就是在pos位置放上皇后
		totalNQueensHelper(n, row+1, col|pos, (leftDiagonal|pos)<<1, (rightDiagonal|pos)>>1, count)
	}
}
