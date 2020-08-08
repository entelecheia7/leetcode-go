package main

import "fmt"

// 198. 打家劫舍
// 你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你 不触动警报装置的情况下 ，一夜之内能够偷窃到的最高金额。
// 提示：
//     0 <= nums.length <= 100
//     0 <= nums[i] <= 400
// https://leetcode-cn.com/problems/house-robber/
func main() {
	// fmt.Println(rob([]int{1, 2, 3, 1}))    // 4
	fmt.Println(rob2([]int{2, 7, 9, 3, 1})) // 12
}

// 法一：动态规划。
// nums[i] 要么纳入统计，要么不纳入统计。
// dp[i]变成二维，dp[i][1]表示偷nums[i]的最大值，dp[i][0]表示不偷nums[i]的最大值
// 这时dp变成二维
// 可以通过改变dp[i]的定义将问题简化，将dp[i]定义为含nums[i]的最大值，则：
// dp[2] = getMax(nums[1]+dp[0], dp[1])
// dp[3] = getMax(nums[2]+dp[1], dp[2])
// 递推公式：dp[n] = getMax(nums[n-1]+ dp[n-2], dp[n-1])
// 空间、时间O(n)
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	state := make([]int, n+1) // state[i]表示前i个元素可以获取的最大结果
	state[0] = 0
	state[1] = nums[0]
	for i := 2; i <= n; i++ {
		state[i] = getMax(nums[i-1]+state[i-2], state[i-1])
	}

	return state[n]
}

// 法二：对法一的递推进行改进，节省空间
func rob2(nums []int) (result int) {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	pp, p := 0, nums[0]
	for i := 1; i < n; i++ {
		result = getMax(nums[i]+pp, p)
		pp, p = p, result
	}
	return result
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
