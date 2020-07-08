package main

import (
	"fmt"
)

// 1. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
// https://leetcode-cn.com/problems/two-sum/
func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

// 使用一个map记录遍历过的值和target的差值
func twoSum(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	state := make(map[int]int, n)
	for i := 0; i < n; i++ {
		if v, exist := state[nums[i]]; exist {
			return []int{v, i}
		}
		state[target-nums[i]] = i
	}

	return nil
}

// 如果是升序数组可以使用双指针
