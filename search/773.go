package main

import (
	"fmt"
)

// 773. 滑动谜题
// 在一个 2 x 3 的板上（board）有 5 块砖瓦，用数字 1~5 来表示, 以及一块空缺用 0 来表示.
// 一次移动定义为选择 0 与一个相邻的数字（上下左右）进行交换.
// 最终当板 board 的结果是 [[1,2,3],[4,5,0]] 谜板被解开。
// 给出一个谜板的初始状态，返回最少可以通过多少次移动解开谜板，如果不能解开谜板，则返回 -1 。
// 提示：
//     board 是一个如上所述的 2 x 3 的数组.
//     board[i][j] 是一个 [0, 1, 2, 3, 4, 5] 的排列.
// https://leetcode-cn.com/problems/sliding-puzzle
func main() {

	fmt.Println(slidingPuzzle([][]int{
		{1, 2, 3},
		{5, 4, 0}})) //-1
	// fmt.Println(slidingPuzzle([][]int{
	// 	{1, 2, 3},
	// 	{4, 0, 5}})) // 1
	// fmt.Println(slidingPuzzle([][]int{
	// 	{4, 1, 2},
	// 	{5, 0, 3}})) // 5
	// fmt.Println(slidingPuzzle([][]int{
	// 	{3, 2, 4},
	// 	{1, 5, 0}})) // 14
}

// 法一：BFS
// 本题可以视为一个图问题，每个数字代表一个顶点，可以交换的顶点用边相连
// 则问题转化为求最短路径问题
func slidingPuzzle(board [][]int) int {

}

// 法二：A*
// func slidingPuzzle2(board [][]int) int {
// }
