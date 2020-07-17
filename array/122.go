package main

import "fmt"

// 122. 买卖股票的最佳时机 II
// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 提示：
//     1 <= prices.length <= 3 * 10 ^ 4
//     0 <= prices[i] <= 10 ^ 4
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))             // 7
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))                // 4
	fmt.Println(maxProfitOptimization([]int{7, 1, 5, 3, 6, 4})) // 7
	fmt.Println(maxProfitOptimization([]int{1, 2, 3, 4, 5}))    // 4
}

// 这道题有一个误区：
// 不应该寻找差值最大的两个点买入卖出，而要寻找价差频繁交易.
// 最大价差策略无法处理回撤再上升的情况。
// 比如 1, 5, 3, 6，买卖两次的利润大于最大价差策略的买卖一次
// 把曲线中的每一个高点（高于相邻的两侧），视为一个卖点
// 而买点是卖点左侧最近的低点（低于相邻两侧）
// 时间复杂度O(n)，空间复杂度O(1)
func maxProfit(prices []int) (profit int) {
	n := len(prices)
	if n == 1 {
		return 0
	}
	buyDay := 0
	for buyDay < n {
		// dertermine a buy day
		for buyDay < n-1 && prices[buyDay] >= prices[buyDay+1] {
			buyDay++
		}
		if buyDay == n-1 {
			break
		}
		sellDay := buyDay + 1 // a certain sell day
		// try to find a higher price day
		for sellDay < n-1 && prices[sellDay] < prices[sellDay+1] {
			sellDay++
		}
		profit += prices[sellDay] - prices[buyDay]
		buyDay = sellDay + 1
	}
	return profit
}

// 对前一个方法进行优化，只遍历一次
// 优化的关键在于：低点-高点相当于一段递增的线段，那么我们只需要计算相邻的递增价格的收益
// best
func maxProfitOptimization(prices []int) (profit int) {
	n := len(prices)
	if n == 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}

	return profit
}
