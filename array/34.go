package main

import (
	"fmt"
)

// 34. 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
// 你的算法时间复杂度必须是 O(log n) 级别。
// 如果数组中不存在目标值，返回 [-1, -1]。
// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
}

// 寻找右侧边界的二分查找
func searchRange(nums []int, target int) (result []int) {
	result = []int{-1, -1}
	if len(nums) == 0 {
		return
	}
	left, right := 0, len(nums)
	// 寻找右侧边界
	for left < right {
		mid := left + ((right - left) >> 1)
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// left-1 即为右侧边界
	if left == 0 || nums[left-1] != target {
		return
	}
	result[1] = left - 1
	result[0] = result[1]
	// search for left border
	for result[0] >= 0 && nums[result[0]] == target {
		result[0] -= 1
	}
	result[0] += 1

	return
}