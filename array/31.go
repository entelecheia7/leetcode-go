package main

import (
	"fmt"
	"sort"
)

// 31. 下一个排列
// 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 必须原地修改，只允许使用额外常数空间。
// 以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
// 1,2,3 → 1,3,2
// 3,2,1 → 1,2,3
// 1,1,5 → 1,5,1
// https://leetcode-cn.com/problems/next-permutation/
func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // 132

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums) // 123

	nums = []int{1, 3, 2}
	nextPermutation(nums)
	fmt.Println(nums) // 213
}

// 法一：O(n)
// 从右侧开始，将大数和它左侧的小数交换位置，即可获得一个更大的数
// 为了满足增幅小的要求：大数和小数要的位置都要尽可能靠右，交换后将大数右侧的数字升序排列
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 从右侧开始，找到nums[i] > nums[i-1]的组合，交换位置
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			break
		}
	}
	// nums[i]右侧的元素是降序排列的
	//
}
