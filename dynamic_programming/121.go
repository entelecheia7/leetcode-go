package main

import (
	"fmt"
	"math"
)

// 121. 买卖股票的最佳时机
// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
// 注意：你不能在买入股票前卖出股票。
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
func main() {
	fmt.Println(maxProfit3([]int{7, 6, 4, 3, 1}))    // 0
	fmt.Println(maxProfit3([]int{7, 1, 5, 3, 6, 4})) // 5
}

// 动态规划
// 法一：dp[i]表示prices[i]右侧大于prices[i]的最大数字，从右向左计算
// 也可以将dp[i]定义为 prices[i]左侧的最小数字，这样是从左向右计算
func maxProfit(prices []int) (profit int) {
	n := len(prices)
	maxRight := make([]int, n+1)
	maxRight[n] = math.MinInt64
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = getMax(maxRight[i+1], prices[i+1])
		if maxRight[i] != math.MinInt64 {
			profit = getMax(profit, maxRight[i]-prices[i])
		}
	}
	return profit
}

// 法二：优化法一的空间复杂度
// best
func maxProfit2(prices []int) (profit int) {
	pre := math.MinInt64
	for i := len(prices) - 2; i >= 0; i-- {
		pre = getMax(pre, prices[i+1])
		if pre != math.MinInt64 {
			profit = getMax(profit, pre-prices[i])
		}
	}
	return profit
}

// 法三：计算递增区间的差价，即为最大利润
// best
func maxProfit3(prices []int) (profit int) {
	curProfit := 0
	for i := 1; i < len(prices); i++ {
		curProfit = getMax(0, curProfit+prices[i]-prices[i-1])
		profit = getMax(profit, curProfit)
	}
	return profit
}

// 法四：单调栈，维护一个递增的栈
func maxProfit4(prices []int) (profit int) {
	n := len(prices)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && stack[len(stack)-1] > prices[i] {
			profit = getMax(stack[len(stack)-1]-stack[0], profit)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, prices[i])
	}
	if len(stack) > 1 {
		return getMax(stack[len(stack)-1]-stack[0], profit)
	}
	return profit
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
