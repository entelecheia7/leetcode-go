package main

import (
	"container/heap"
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
	fmt.Println(slidingPuzzle2([][]int{
		{1, 2, 3},
		{5, 4, 0}})) //-1
	fmt.Println(slidingPuzzle2([][]int{
		{1, 2, 3},
		{4, 0, 5}})) // 1
	fmt.Println(slidingPuzzle2([][]int{
		{4, 1, 2},
		{5, 0, 3}})) // 5
	fmt.Println(slidingPuzzle2([][]int{
		{3, 2, 4},
		{1, 5, 0}})) // 14
}

// 法一：BFS
// 本题可以视为一个图问题，每个数字代表一个顶点，可以交换的顶点用边相连
// 则问题转化为求最短路径问题
// move代表将board的六个点标记为数字，0在key位置可以移动到的位置
var move = [6][]int{0: {1, 3}, 1: {0, 2, 4}, 2: {1, 5}, 3: {0, 4}, 4: {1, 3, 5}, 5: {2, 4}}

// 将board转换为string，二维变一维
// 再按照可变换的序列进行变换
func slidingPuzzle(board [][]int) int {
	orig := make([]byte, 6)
	zero := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				zero = i*3 + j
			}
			orig[i*3+j] = byte('1' + (board[i][j] - 1))
		}
	}
	queue := []element{{orig, zero}}
	visited := make(map[string]bool)
	visited[string(orig)] = true
	step := 0
	for len(queue) > 0 {
		size := len(queue)
		newQ := []element{}
		for i := 0; i < size; i++ {
			cur := queue[i]
			if string(cur.board) == "123450" {
				return step
			}
			for _, next := range move[cur.zero] {
				tmp := make([]byte, 6)
				copy(tmp, cur.board)
				tmp[cur.zero], tmp[next] = tmp[next], tmp[cur.zero]
				if !visited[string(tmp)] {
					newQ = append(newQ, element{tmp, next})
					visited[string(tmp)] = true
				}
			}
		}
		step++
		queue = newQ
	}

	return -1
}

type element struct {
	board []byte
	zero  int
}

// 法二：A*
func slidingPuzzle2(board [][]int) int {
	// 将矩阵转换为 string
	start := make([]byte, 0, 6)
	zero := 0
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			start = append(start, byte(board[i][j]+'0'))
			if board[i][j] == 0 {
				zero = 3*i + j
			}
		}
	}
	var pq priorityQueue
	pq = []priorityElement{{start, zero, 0, heuristicHelper(start)}}
	visited := make(map[string]bool)

	for len(pq) > 0 {
		cur := heap.Pop(&pq).(priorityElement)
		if string(cur.board) == "123450" {
			return cur.dist
		}
		if visited[string(cur.board)] {
			continue
		}
		for _, next := range move[cur.zero] {
			tmp := make([]byte, 6)
			copy(tmp, cur.board)
			tmp[cur.zero], tmp[next] = tmp[next], tmp[cur.zero]
			if visited[string(tmp)] {
				continue
			}
			heap.Push(&pq, priorityElement{
				board:    tmp,
				zero:     next,
				dist:     cur.dist + 1,
				priority: heuristicHelper(tmp) + cur.dist + 1})
		}
		visited[string(cur.board)] = true
	}
	return -1
}

type priorityQueue []priorityElement

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq *priorityQueue) Swap(i, j int)     { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *priorityQueue) Pop() interface{} {
	last := (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	return last
}
func (pq *priorityQueue) Push(x interface{}) { (*pq) = append((*pq), x.(priorityElement)) }

type priorityElement struct {
	board    []byte
	zero     int
	dist     int
	priority int
}

// 估价函数计算的是每个点和终点的坐标距离的差的和
// 值越小越好
var goal = [6][2]int{0: {1, 2}, 1: {0, 0}, 2: {0, 1}, 3: {0, 2}, 4: {1, 0}, 5: {1, 1}}

// 字符串字符位置对应的矩阵坐标位置
var pos = [6][2]int{0: {0, 0}, 1: {0, 1}, 2: {0, 2}, 3: {1, 0}, 4: {1, 1}, 5: {1, 2}}

func heuristicHelper(board []byte) (priority int) {
	for i := 0; i < 6; i++ {
		n := board[i] - '0' // goal[n]为该位置字符的目标位置
		priority += getAbs(pos[i][0]-goal[n][0]) + getAbs(pos[i][1]-goal[n][1])
	}
	return
}

func getAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
