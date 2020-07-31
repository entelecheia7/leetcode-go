package main

import (
	"fmt"
)

// 面试题 08.03. 魔术索引
// 魔术索引。 在数组A[0...n-1]中，有所谓的魔术索引，满足条件A[i] = i。给定一个有序整数数组，编写一种方法找出魔术索引，若有的话，在数组A中找出一个魔术索引，如果没有，则返回-1。若有多个魔术索引，返回索引值最小的一个。
// https://leetcode-cn.com/problems/magic-index-lcci/
func main() {
	fmt.Println(findMagicIndex([]int{0, 2, 3, 4, 5}))
}

// 法一：暴力搜索，略

// 法二：二分查找
func findMagicIndex(nums []int) int {
	var helper func(int, int) int
	helper = func(left, right int) int {
		if left > right {
			return -1
		}
		mid := left + ((right - left) >> 1)
		leftResult := helper(left, mid-1)
		if leftResult != -1 {
			return leftResult
		} else if mid == nums[mid] {
			return mid
		}
		return helper(mid+1, right)
	}
	return helper(0, len(nums)-1)
}
