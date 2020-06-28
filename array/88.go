package main

import "fmt"

// 88. 合并两个有序数组
// 给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
// 说明:
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
// https://leetcode-cn.com/problems/merge-sorted-array
func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// merge(nums1, 3, []int{2, 5, 6}, 3)
	// nums1 := []int{1}
	// merge(nums1, 1, []int{}, 0)
	nums1 := []int{4, 0, 0, 0, 0, 0}
	merge(nums1, 1, []int{1, 2, 3, 5, 6}, 5)
	// nums1 := []int{4, 5, 6, 0, 0, 0}
	// merge(nums1, 3, []int{1, 2, 3}, 3)
	fmt.Println(nums1)
}

// 思路：利用双指针，从m+n-1的位置开始填充两个数组中更大的那个元素
func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[pos] = nums1[m]
			m--
		} else {
			nums1[pos] = nums2[n]
			n--
		}
		pos--
	}
	for ; n >= 0; n-- {
		nums1[pos] = nums2[n]
		pos--
	}
}
