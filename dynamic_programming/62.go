package main

import (
	"fmt"
)

// 62. 不同路径
// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
// 问总共有多少条不同的路径？
// 提示：
//     1 <= m, n <= 100
//     题目数据保证答案小于等于 2 * 10 ^ 9
// https://leetcode-cn.com/problems/unique-paths/
func main() {
	fmt.Println(uniquePaths2(3, 2)) // 3
	fmt.Println(uniquePaths2(7, 3)) // 28
}

// 动态规划
// 法一：dp[i][j]表示在网格中从起点到某位置时的不同路径数
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
// O(mn)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 法二：对法一空间复杂度的优化，优化为 O(n)
// best
func uniquePaths2(m int, n int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j > 0 {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
