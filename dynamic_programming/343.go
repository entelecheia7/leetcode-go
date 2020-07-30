package main

import "fmt"

// 343. 整数拆分
// 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
// 说明: 你可以假设 n 不小于 2 且不大于 58。
// https://leetcode-cn.com/problems/integer-break/
func main() {
	fmt.Println(integerBreak(10)) // 36=3*3*4
}

// 动态规划
// dp[i]表示 i 的最大乘积
// 动态转移方程：dp[i] = getMax(j *(i-j), j*dp[j-1]) (1<=j<i)
// O(n^2)
func integerBreak(n int) int {
	if n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = getMax(getMax(j*(i-j), dp[i]), j*dp[i-j])
		}
	}
	return dp[n]
}

// 根据数学方法推导可以优化至O(n)，甚至O(1)
// 见：https://leetcode-cn.com/problems/integer-break/solution/343-zheng-shu-chai-fen-tan-xin-by-jyd/

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
