package main

import (
	"fmt"
)

// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// https://leetcode-cn.com/problems/subsets/
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

// 法一：暴力法，逐个生成长度为1、2、3……len(nums)的子集，复杂度过高，略

// 法二：回溯
func subsets(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {
		subsetBTHelper(nums, 0, []int{}, i+1, &result)
	}

	return result
}
func subsetBTHelper(nums []int, from int, cur []int, length int, result *[][]int) {
	if len(cur) == length {
		tmp := make([]int, length)
		copy(tmp, cur)
		*result = append(*result, tmp)
	}
	for i := from; i < len(nums); i++ {
		subsetBTHelper(nums, i+1, append(cur, nums[i]), length, result)
	}
}
