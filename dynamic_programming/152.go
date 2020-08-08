package main

import (
	"fmt"
)

// 152. 乘积最大子数组
// 给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
// https://leetcode-cn.com/problems/maximum-product-subarray/
func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4})) // 6
	fmt.Println(maxProduct([]int{-2, 0, -1}))   // 0
	fmt.Println(maxProduct([]int{-3, -1, -1}))  // 3
	fmt.Println(maxProduct([]int{-2, 3, -4}))   // 24
}

// dp[i]以nums[i]结尾的最大连续子数组乘积
// 如果只存在非负数，则 dp[i] = max(nums[i], dp[i-1]*nums[i])
// 当负数出现时，需要求以 nums[i]结尾的最小乘积
// 因此需要保留最小和最大两个dp数组
// 由于 dp[i]只和 dp[i-1]有关，可以只使用O(1)空间
func maxProduct(nums []int) (max int) {
	if len(nums) == 0 {
		return
	}
	n := len(nums)
	preMax, preMin := 1, 1
	max = nums[0]
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			preMax, preMin = preMin, preMax
		}
		preMax = getMax(nums[i], preMax*nums[i])
		preMin = getMin(nums[i], preMin*nums[i])
		max = getMax(preMax, max)
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
