package main

import (
	"fmt"
)

// 213. 打家劫舍 II
// 你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都围成一圈，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。
// 给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。
// https://leetcode-cn.com/problems/house-robber-ii/
func main() {
	fmt.Println(rob([]int{2, 3, 2}))    // 3
	fmt.Println(rob([]int{1, 2, 3, 1})) // 4
}

// 动态规划
// 分两种情况：从第一间偷至到倒数第二间、从第二间偷至最后一间
func rob(nums []int) (result int) {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	return getMax(robHelper(nums[:n-1]), robHelper(nums[1:]))
}
func robHelper(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	pp, p := 0, nums[0]
	var cur int
	for i := 1; i < n; i++ {
		cur = getMax(nums[i]+pp, p)
		pp, p = p, cur
	}
	return p
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
