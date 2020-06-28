package main

import (
	"fmt"
)

// 283. 移动零
// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 说明:
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
// https://leetcode-cn.com/problems/move-zeroes/
func main() {
	nums := []int{0, 1, 0, 3, 12}
	// best
	moveZeroes(nums)
	fmt.Println(nums)
}

// 双指针：快指针负责遍历数组，慢指针记录数组前端非0序列的下一个位置
// 当遍历到非0元素时，就和慢指针位置交换
func moveZeroes1(nums []int) {
	nextPositivePos := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 && i > nextPositivePos {
			nums[nextPositivePos], nums[i] = nums[i], nums[nextPositivePos]
			nextPositivePos++
		}
	}
}

// 另一种写法
func moveZeroes2(nums []int) {
	nextPositivePos := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[nextPositivePos] = nums[i]
			nextPositivePos++
		}
	}
	for ; nextPositivePos < n; nextPositivePos++ {
		nums[nextPositivePos] = 0
	}
}

// 双指针优化
func moveZeroes(nums []int) {
	nextPositivePos := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			if i > nextPositivePos {
				nums[nextPositivePos], nums[i] = nums[i], nums[nextPositivePos]
				nextPositivePos++
			} else {
				nextPositivePos++
			}
		}
	}
}
