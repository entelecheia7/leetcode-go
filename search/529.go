package main

import "fmt"

// 529. 扫雷游戏
// 让我们一起来玩扫雷游戏！
// 给定一个代表游戏板的二维字符矩阵。 'M' 代表一个未挖出的地雷，'E' 代表一个未挖出的空方块，'B' 代表没有相邻（上，下，左，右，和所有4个对角线）地雷的已挖出的空白方块，数字（'1' 到 '8'）表示有多少地雷与这块已挖出的方块相邻，'X' 则表示一个已挖出的地雷。
// 现在给出在所有未挖出的方块中（'M'或者'E'）的下一个点击位置（行和列索引），根据以下规则，返回相应位置被点击后对应的面板：
//     如果一个地雷（'M'）被挖出，游戏就结束了- 把它改为 'X'。
//     如果一个没有相邻地雷的空方块（'E'）被挖出，修改它为（'B'），并且所有和其相邻的方块都应该被递归地揭露。
//     如果一个至少与一个地雷相邻的空方块（'E'）被挖出，修改它为数字（'1'到'8'），表示相邻地雷的数量。
//     如果在此次点击中，若无更多方块可被揭露，则返回面板。
// 示例 1：
// 输入:
// [['E', 'E', 'E', 'E', 'E'],
//  ['E', 'E', 'M', 'E', 'E'],
//  ['E', 'E', 'E', 'E', 'E'],
//  ['E', 'E', 'E', 'E', 'E']]
// Click : [3,0]
// 输出:
// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'M', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]
// 示例 2：
// 输入:
// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'M', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]
// Click : [1,2]
// 输出:
// [['B', '1', 'E', '1', 'B'],
//  ['B', '1', 'X', '1', 'B'],
//  ['B', '1', '1', '1', 'B'],
//  ['B', 'B', 'B', 'B', 'B']]
// 注意：
//     输入矩阵的宽和高的范围为 [1,50]。
//     点击的位置只能是未被挖出的方块 ('M' 或者 'E')，这也意味着面板至少包含一个可点击的方块。
//     输入面板不会是游戏结束的状态（即有地雷已被挖出）。
//     简单起见，未提及的规则在这个问题中可被忽略。例如，当游戏结束时你不需要挖出所有地雷，考虑所有你可能赢得游戏或标记方块的情况。
// https://leetcode-cn.com/problems/minesweeper/
func main() {
	fmt.Println(updateBoard([][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}},
		[]int{3, 0}))
}

// 法一：BFS
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	lx := len(board)
	ly := len(board[0])
	queue := [][2]int{{click[0], click[1]}}
	for len(queue) > 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]

		relative := [][2]int{}
		var mine byte = '0' // 地雷数量
		if x >= 1 {         // top line
			if y >= 1 {
				if board[x-1][y-1] == 'M' {
					mine++
				} else if board[x-1][y-1] == 'E' {
					relative = append(relative, [2]int{x - 1, y - 1})
				}
			}
			if y+1 < ly {
				if board[x-1][y+1] == 'M' {
					mine++
				} else if board[x-1][y+1] == 'E' {
					relative = append(relative, [2]int{x - 1, y + 1})
				}
			}
			if board[x-1][y] == 'M' {
				mine++
			} else if board[x-1][y] == 'E' {
				relative = append(relative, [2]int{x - 1, y})
			}
		}
		if x+1 < lx { //bottom line
			if board[x+1][y] == 'M' {
				mine++
			} else if board[x+1][y] == 'E' {
				relative = append(relative, [2]int{x + 1, y})
			}
			if y >= 1 {
				if board[x+1][y-1] == 'M' {
					mine++
				} else if board[x+1][y-1] == 'E' {
					relative = append(relative, [2]int{x + 1, y - 1})
				}
			}
			if y+1 < ly {
				if board[x+1][y+1] == 'M' {
					mine++
				} else if board[x+1][y+1] == 'E' {
					relative = append(relative, [2]int{x + 1, y + 1})
				}
			}
		}
		// middle line
		if y >= 1 {
			if board[x][y-1] == 'M' {
				mine++
			} else if board[x][y-1] == 'E' {
				relative = append(relative, [2]int{x, y - 1})
			}
		}
		if y+1 < ly {
			if board[x][y+1] == 'M' {
				mine++
			} else if board[x][y+1] == 'E' {
				relative = append(relative, [2]int{x, y + 1})
			}
		}
		if mine == '0' {
			board[x][y] = 'B'
			queue = append(queue, relative...)
		} else {
			board[x][y] = mine
		}

		printBoard(board)
	}

	return board
}

func printBoard(board [][]byte) {
	s := ""
	for i := 0; i < len(board); i++ {
		s += string(board[i]) + "\n"
	}
	fmt.Println(s)
}
