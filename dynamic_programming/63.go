package main

import (
	"fmt"
)

// 63. 不同路径 II
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
// 说明：m 和 n 的值均不超过 100。
// https://leetcode-cn.com/problems/unique-paths-ii/
func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})) // 2
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1, 0},
	})) // 0
}

// 动态规划
// 时间 O(mn)，空间 O(n)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1
	for _, row := range obstacleGrid {
		for j := 0; j < n; j++ {
			if row[j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
