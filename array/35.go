package main

import "fmt"

// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 你可以假设数组中无重复元素。
// https://leetcode-cn.com/problems/search-insert-position/
func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) //2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) //1
}

// 二分查找
// 当数组中不存在target，要找出大于target的第一个元素
func searchInsert(nums []int, target int) (result int) {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
