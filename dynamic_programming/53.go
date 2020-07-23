package main

import (
	"fmt"
)

// 53. 最大子序和
// 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 进阶:
// 如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
// https://leetcode-cn.com/problems/maximum-subarray/
func main() {
	// fmt.Println(maxSubArray3([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray3([]int{-2, 1})) // 1
}

// 法一：动态规划
// dp[i] 表示以nums[i]结尾的最大子序和
func maxSubArray(nums []int) (max int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = getMax(dp[i-1]+nums[i], nums[i])
		max = getMax(max, dp[i])
	}
	return max
}

// 法二：对法一的优化
// best
func maxSubArray2(nums []int) (max int) {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max = nums[0]
	pre := nums[0]
	for i := 1; i < n; i++ {
		pre = getMax(nums[i], pre+nums[i])
		max = getMax(max, pre)
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 法三：分治
// O(nlogn)
func maxSubArray3(nums []int) (max int) {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	mid := n >> 1
	maxLeft := maxSubArray3(nums[0:mid])
	maxRight := maxSubArray3(nums[mid:])
	// 计算中间向两边最大子序和
	l := nums[mid-1] // 向左侧最大值
	tmp := 0
	for i := mid - 1; i >= 0; i-- {
		tmp += nums[i]
		l = getMax(tmp, l)
	}
	r := nums[mid] // 向右侧最大值
	tmp = 0
	for i := mid; i < n; i++ {
		tmp += nums[i]
		r = getMax(tmp, r)
	}

	return getMax(getMax(maxLeft, maxRight), l+r)
}
