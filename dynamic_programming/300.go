package main

import (
	"fmt"
)

// 300. 最长上升子序列
// 给定一个无序的整数数组，找到其中最长上升子序列的长度。
// 说明:
//     可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
//     你算法的时间复杂度应该为 O(n2) 。
// 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
// https://leetcode-cn.com/problems/longest-increasing-subsequence
func main() {
	fmt.Println(lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18})) // 4
	fmt.Println(lengthOfLIS2([]int{1, 3, 6, 7, 9, 4, 10, 5, 6})) // 6
}

// 法一：动态规划
// dp[i]代表nums[i]结尾的最长上升子序列长度
// dp[i]的取决于dp[j]，0<=j<i 且 nums[j] < nums[i]，则dp[i] = dp[j]+1
// 时间 O(n^2)，空间 O(n)
func lengthOfLIS(nums []int) (result int) {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = getMax(dp[i], dp[j]+1)
			}
		}
		result = getMax(result, dp[i])
	}
	return result
}

// 法二：tail[i]代表长度为i+1的上升子序列中的末位最小值
func lengthOfLIS2(nums []int) (result int) {
	n := len(nums)
	if n <= 1 {
		return n
	}
	tail := make([]int, n)
	tail[0] = nums[0]
	key := 0 // 目前计算完毕的tail[i]的索引
	for i := 1; i < n; i++ {
		if nums[i] > tail[key] {
			key++
			tail[key] = nums[i]
		} else {
			// 在计算完毕的tail范围内查找大于等于nums[i]的最小元素
			// 试图缩小计算完毕的长度为m的子序列的范围
			left, right := 0, key
			for left < right {
				mid := left + ((right - left) >> 1)
				if nums[mid] >= nums[i] {
					right = mid
				} else {
					left = mid + 1
				}
			}
			tail[left] = nums[i]
		}
	}
	return key + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
