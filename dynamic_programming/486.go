package main

import (
	"fmt"
)

// 486. 预测赢家
// 给定一个表示分数的非负整数数组。 玩家 1 从数组任意一端拿取一个分数，随后玩家 2 继续从剩余数组任意一端拿取分数，然后玩家 1 拿，…… 。每次一个玩家只能拿取一个分数，分数被拿取之后不再可取。直到没有剩余分数可取时游戏结束。最终获得分数总和最多的玩家获胜。
// 给定一个表示分数的数组，预测玩家1是否会成为赢家。你可以假设每个玩家的玩法都会使他的分数最大化。
// 提示：
//     1 <= 给定的数组长度 <= 20.
//     数组里所有分数都为非负数且不会大于 10000000 。
//     如果最终两个玩家的分数相等，那么玩家 1 仍为赢家。
// https://leetcode-cn.com/problems/predict-the-winner
func main() {
	fmt.Println(PredictTheWinner2([]int{1, 5, 2}))      // false
	fmt.Println(PredictTheWinner2([]int{1, 5, 233, 7})) // true
}

// 法一：动态规划
// dp[i][j]表示nums[i]到nums[j]范围内，先手的玩家比后手玩家多的分数(考虑先拿nums[i]、nums[j]两种情况，取更大值)
// dp[i][j]如果先取nums[i]，则dp[i][j] = nums[i]-dp[i+1][j]
// dp[i][j]如果先取nums[j]，则dp[i][j] = nums[j]-dp[i][j-1]
// 因此 dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
func PredictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for k := range dp {
		dp[k] = make([]int, n)
		dp[k][k] = nums[k]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = getMax(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][n-1] >= 0
}

// 法二：法一的空间优化
func PredictTheWinner2(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[j] = nums[j]
			} else {
				dp[j] = getMax(nums[i]-dp[j], nums[j]-dp[j-1])
			}
		}
	}

	return dp[n-1] >= 0
}
func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
