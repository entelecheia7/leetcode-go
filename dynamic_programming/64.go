package main

import "fmt"

// 64. 最小路径和
// 给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
// 说明：每次只能向下或者向右移动一步。
// https://leetcode-cn.com/problems/minimum-path-sum/
func main() {
	fmt.Println(minPathSum2([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})) //7
}

// 法一：动态规划
// 每个坐标只能从左侧或上边移动而来
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			dp[i][j] = grid[i][j]
			if i == 0 {
				dp[i][j] += dp[i][j-1]
			} else if j == 0 {
				dp[i][j] += dp[i-1][j]
			} else {
				dp[i][j] += getMin(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[m-1][n-1]
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：动态规划空间复杂度优化
func minPathSum2(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = grid[i][j] + dp[j]
			} else {
				dp[j] = getMin(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}
	return dp[n-1]
}
