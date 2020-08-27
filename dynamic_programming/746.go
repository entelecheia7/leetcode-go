package main

import "fmt"

// 746. 使用最小花费爬楼梯
// 数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。
// 每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。
// 您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。
// 注意：
//     cost 的长度将会在 [2, 1000]。
//     每一个 cost[i] 将会是一个Integer类型，范围为 [0, 999]。
// https://leetcode-cn.com/problems/min-cost-climbing-stairs/
func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))                         // 15
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})) // 6
}

// 法一：动态规划
// dp[i]表示到达第i层的最低花费，dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < n; i++ {
		dp[i] = getMin(dp[i-1], dp[i-2]) + cost[i]
	}
	return getMin(dp[n-2], dp[n-1])
}

// 法二：dp空间优化
func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	pp, p := cost[0], cost[1]
	for i := 2; i < n; i++ {
		cur := getMin(pp, p) + cost[i]
		pp, p = p, cur
	}
	return getMin(pp, p)
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
