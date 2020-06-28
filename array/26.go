package main

import (
	"fmt"
)

// 26. 删除排序数组中的重复项
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := removeDuplicates(nums)
	fmt.Println(nums, n)
}

// 将唯一元素移动到数组开头
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	pos := 1 // 下一个唯一元素放置的位置
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
