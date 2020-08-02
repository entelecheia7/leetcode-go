package main

import (
	"fmt"
)

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 链接：https://leetcode-cn.com/problems/climbing-stairs
func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs2(3))
}

// 递推公式：f(n) = f(n-1)+f(n-2)
// 实现方式：递归、循环
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	pp := 1 // f(n-2)
	p := 2  // f(n-1)
	for i := 3; i <= n; i++ {
		p, pp = p+pp, p
	}
	return p
}

// 变形问题：假设每次可以走1、2、3步，同时相邻两次步数不能相同
// dp[i][j] 代表从上一层台阶走 j 步走到第 i 级台阶时的方法数
// dp[i][j] = dp[]
func climbStairs2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([][]int, n+1)
	for k := range dp {
		dp[k] = make([]int, 4)
	}
	dp[1][1] = 1
	dp[2][2] = 1
	dp[3][3] = 1
	dp[3][2] = 1
	dp[3][1] = 1
	for i := 4; i <= n; i++ {
		dp[i][3] = dp[i-3][2] + dp[i-3][1]
		dp[i][2] = dp[i-2][1] + dp[i-2][3]
		dp[i][1] = dp[i-1][2] + dp[i-1][3]
	}

	return dp[n][1] + dp[n][2] + dp[n][3]
}
