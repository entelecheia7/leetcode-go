package main

import (
	"fmt"
	// "sort"
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

// 从右侧开始，将大数和它左侧的小数交换位置，即可获得一个更大的数
// 为了满足增幅小的要求：大数和小数要的位置都要尽可能靠右，交换后将大数位置右侧的数字升序排列
// O(n)
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 从右侧开始，找到 nums[min] < nums[min+1] 的组合，nums[min]就是那个小数
	min := n - 2
	for ; min >= 0; min-- {
		if nums[min] < nums[min+1] {
			break
		}
	}
	// 如果nums是降序序列，是没有比它更大的排列的，直接reverse整个数组
	// 如果nums不是降序序列，在nums[min+1:]范围内找到：nums[min]<nums[k]，k是大数位置
	if min >= 0 {
		k := n - 1
		for nums[k] <= nums[min] {
			k--
		}
		nums[k], nums[min] = nums[min], nums[k]
	}

	// reverse nums[min+1:]
	for x, y := min+1, n-1; x < y; x, y = x+1, y-1 {
		nums[x], nums[y] = nums[y], nums[x]
	}
}
