package main

import (
	"fmt"
	"math"
)

// 33. 搜索旋转排序数组
// 假设按照升序排序的数组在预先未知的某个点上进行了旋转。
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
// 搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
// 你可以假设数组中不存在重复的元素。
// 你的算法时间复杂度必须是 O(log n) 级别。
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func main() {
	fmt.Println(search2([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}

// 法一：直接进行二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] { // 左侧为有序数组
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右侧为有序数组
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 法二：法一代码的另一种写法
func search2(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		}
		num := nums[mid]
		// 左侧为无序数组，且target<nums[0]
		// 左侧为有序数组，且target>=nums[0]
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			num = nums[mid]
		} else {
			if target < nums[0] {
				num = math.MinInt64
			} else {
				num = math.MaxInt64
			}
		}
		if num < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
