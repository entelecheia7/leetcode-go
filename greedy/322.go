package main

import (
	"fmt"
	"math"
	"sort"
)

// 322. 零钱兑换
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
// 说明:
// 你可以认为每种硬币的数量是无限的。
// https://leetcode-cn.com/problems/coin-change/
func main() {
	// fmt.Println(coinChange([]int{1, 2, 5}, 11)) // 3
	fmt.Println(coinChange([]int{2, 5, 17}, 40)) // 5
	fmt.Println(coinChange([]int{1, 7, 10}, 14)) // 2
}

// 法一：贪心+回溯
// 注意，这道题不能只使用贪心算法，因为局部最优不一定能形成一个合理的组合
// 大硬币的数量过多导致无法形成组合，就减少硬币数量再尝试
// 同时，贪心优先找出的不一定是最优解
func coinChange(coins []int, amount int) (count int) {
	if amount < 0 {
		return -1
	}
	if len(coins) == 0 && amount == 0 {
		return 0
	}
	sort.Ints(coins)
	count = math.MaxInt64
	coinChangeBacktracingHelper(coins, amount, 0, &count)
	if count == math.MaxInt64 {
		return -1
	}
	return count
}

// coins是有序数组，amount>=0
func coinChangeBacktracingHelper(coins []int, amount int, curCount int, minCount *int) {
	if amount == 0 {
		*minCount = getMin(*minCount, curCount)
		return
	} else if len(coins) == 0 || amount < coins[0] { // 无法得出有效组合
		return
	}
	// 用最大币值尝试
	n := len(coins)
	maxCoin := coins[n-1]
	coins = coins[:n-1]
	num := amount / maxCoin // 最多使用num个maxCoin硬币，最少使用0个
	for j := num; j >= 0 && j+curCount < *minCount; j-- {
		coinChangeBacktracingHelper(coins, amount-j*maxCoin, curCount+j, minCount)
	}
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：动态规划
func coinChange2(coins []int, amount int) (count int) {
}
