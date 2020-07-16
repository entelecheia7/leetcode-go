package main

import (
	"fmt"
	"math"
)

// 120. 三角形最小路径和
// 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
// 相邻的结点 在这里指的是 下标 与 上一层结点下标 相同或者等于 上一层结点下标 + 1 的两个结点。
// https://leetcode-cn.com/problems/triangle/
func main() {
	fmt.Println(minimumTotal2([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

// 法一：动态规划，O(n^2)
func minimumTotal(triangle [][]int) (minPath int) {
	level := len(triangle)
	if level == 0 || len(triangle[0]) == 0 {
		return 0
	}
	if level == 1 {
		return triangle[0][0]
	}
	state := make([][]int, level)
	for i := 0; i < level; i++ {
		state[i] = make([]int, i+1)
	}
	state[0][0] = triangle[0][0]
	for i := 1; i < level; i++ { // 计算每行的路径
		for j := 0; j <= i; j++ {
			state[i][j] = triangle[i][j]
			if i == j {
				state[i][j] += state[i-1][j-1]
			} else if j == 0 {
				state[i][j] += state[i-1][j]
			} else {
				state[i][j] += getMin(state[i-1][j-1], state[i-1][j])
			}
		}
	}
	minPath = math.MaxInt64
	for _, path := range state[level-1] {
		minPath = getMin(minPath, path)
	}
	return minPath
}

// 法二：对法一进行空间优化，只使用一个一维数组保存中间状态
// 同时自底向上计算，省去正向计算最后查找最小路径的一次循环
// 时间O(n^2)，空间O(n)，n是triangle行数
// best
func minimumTotal2(triangle [][]int) int {
	level := len(triangle)
	if level == 0 || len(triangle[0]) == 0 {
		return 0
	}
	if level == 1 {
		return triangle[0][0]
	}
	state := make([]int, level+1)
	for i := level - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			state[j] = triangle[i][j] + getMin(state[j], state[j+1])
		}
	}

	return state[0]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
