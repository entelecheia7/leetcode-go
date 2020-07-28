package main

import "fmt"

// 221. 最大正方形
// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
// https://leetcode-cn.com/problems/maximal-square/
func main() {
	// fmt.Println(maximalSquare2([][]byte{
	// 	{'1', '0', '1', '0', '0'},
	// 	{'1', '0', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '0', '0', '1', '0'}})) //4
	// fmt.Println(maximalSquare2([][]byte{
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '1', '1', '0'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'1', '1', '1', '1', '1'},
	// 	{'0', '0', '1', '1', '1'}})) //16

	// fmt.Println(maximalSquare2([][]byte{
	// 	{'0', '0', '0', '1'},
	// 	{'1', '1', '0', '1'},
	// 	{'1', '1', '1', '1'},
	// 	{'0', '1', '1', '1'},
	// 	{'0', '1', '1', '1'}})) //9

	fmt.Println(maximalSquare2([][]byte{
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'}})) // 4
}

// 法一：动态规划
// dp[i][j]表示以matrix[i][j]为正方形右下角的最大正方形边长
// 如果 matrix[i][j] 是 1，dp[i][j] = min(dp[i-1][j], dp[i-1][j-1], dp[i][j-1])+1
func maximalSquare(matrix [][]byte) (maxArea int) {
	// check
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for k := range dp {
		dp[k] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				if i > 0 && j > 0 {
					dp[i][j] = getMin(getMin(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
				}
				maxArea = getMax(maxArea, dp[i][j]*dp[i][j])
			}
		}
	}

	return maxArea
}

// 法二：优化法一的空间
// best
func maximalSquare2(matrix [][]byte) int {
	// check
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	maxBorder := 0
	dp := make([]int, n)
	var leftup int
	for i := 0; i < m; i++ {
		leftup = 0
		for j := 0; j < n; j++ {
			nextLeftup := dp[j]
			if matrix[i][j] == '1' {
				if maxBorder == 0 {
					maxBorder = 1
				}
				if i == 0 || j == 0 {
					dp[j] = 1
				} else {
					dp[j] = getMin(leftup, getMin(dp[j], dp[j-1])) + 1
				}
				if dp[j] > maxBorder {
					maxBorder = dp[j]
				}
			} else {
				dp[j] = 0
			}

			leftup = nextLeftup
		}
	}

	return maxBorder * maxBorder
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
