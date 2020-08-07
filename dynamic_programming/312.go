package main

import "fmt"

// 312. 戳气球
// 有 n 个气球，编号为0 到 n-1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
// 现在要求你戳破所有的气球。如果你戳破气球 i ，就可以获得 nums[left] * nums[i] * nums[right] 个硬币。 这里的 left 和 right 代表和 i 相邻的两个气球的序号。注意当你戳破了气球 i 后，气球 left 和气球 right 就变成了相邻的气球。
// 求所能获得硬币的最大数量。
// 说明:
//     你可以假设 nums[-1] = nums[n] = 1，但注意它们不是真实存在的所以并不能被戳破。
//     0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100
// https://leetcode-cn.com/problems/burst-balloons/
func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8})) //167
}

// 法一：回溯，时间复杂度高，略

// 法二：动态规划
// dp[i][j] 表示，戳破气球 i 和气球 j 之间（不含 i、 j）的所有气球，可获得的最高分数。
// 设 nums (i, j)间最后一个被戳的气球是 k，则 dp[i][j] = dp[i][k] + nums[k]*nums[i]*nums[j]+dp[k][j]
// 固定 i，j，对 k 进行枚举
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	balls := make([]int, n+2)
	balls[0] = 1
	copy(balls[1:], nums)
	balls[n+1] = 1

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	// 对于任意的 dp[i][j]，需要满足 dp[i][k] 和 dp[k][j]已被计算
	// i < k < j
	// 因此 i 需要倒序计算，j从左到右计算
	for i := n + 1; i >= 0; i-- {
		for j := i + 2; j < n+2; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = getMax(dp[i][k]+balls[k]*balls[i]*balls[j]+dp[k][j], dp[i][j])
			}
		}
	}

	return dp[0][n+1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
