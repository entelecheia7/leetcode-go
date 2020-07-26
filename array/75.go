package main

import (
	"fmt"
)

// 75. 颜色分类
// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
// 注意:
// 不能使用代码库中的排序函数来解决这道题。
// 进阶：
//     一个直观的解决方案是使用计数排序的两趟扫描算法。
//     首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
//     你能想出一个仅使用常数空间的一趟扫描算法吗？
// https://leetcode-cn.com/problems/sort-colors
func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	nums = []int{0, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

// 法一：双指针，O(n)
func sortColors(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	l, r := 0, n-1
	for i := 0; i <= r; i++ {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
			i--
		}
	}
}

// 法二：排序算法，略
