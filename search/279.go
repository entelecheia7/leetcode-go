package main

import "fmt"

// 279. 完全平方数
// 给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。
// https://leetcode-cn.com/problems/perfect-squares/
func main() {
	fmt.Println(numSquares(3))
}

// 法一：动态规划
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		for x := 2; i-x*x >= 0; x++ {
			dp[i] = getMin(dp[i], dp[i-x*x]+1)
		}
	}
	return dp[n]
}

// 法二：BFS
// 将n视为根节点，减去一个平方数后的值视为下一层
// 问题转化为求树的最小层数
func numSquares2(n int) (level int) {
	if n == 1 {
		return 1
	}
	queue := []int{n}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			cur := queue[i]
			for j := 1; j*j <= cur; j++ {
				next := cur - j*j
				if next == 0 {
					return level
				}
				if !visited[next] {
					queue = append(queue, next)
					visited[next] = true
				}
			}
		}
	}
	return level
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
