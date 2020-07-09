package main

import (
	"fmt"
)

// 350. 两个数组的交集 II
// 给定两个数组，编写一个函数来计算它们的交集。说明：
//     输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//     我们可以不考虑输出结果的顺序。
// 进阶:
//     如果给定的数组已经排好序呢？你将如何优化你的算法？
//     如果 nums1 的大小比 nums2 小很多，哪种方法更优？
//     如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

// 使用一个额外的map进行统计，两遍遍历，O(n)
func intersect(nums1 []int, nums2 []int) (result []int) {
	if nums1 == nil || nums2 == nil {
		return nil
	}
	m := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if v, ok := m[num]; ok && v != 0 {
			result = append(result, num)
			m[num]--
		}
	}
	return result
}

// 进阶：如果数组已排序，双指针，一遍遍历
func intersect2(nums1 []int, nums2 []int) (result []int) {
	if nums1 == nil || nums2 == nil {
		return nil
	}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}

// 进阶：如果nums1比nums2小很多，使用小的数组遍历生成map

// 进阶：如果磁盘空间有限，将两个数组排序，再双指针遍历比较
