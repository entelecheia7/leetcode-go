package main

import (
	"fmt"
)

// 349. 两个数组的交集
// 给定两个数组，编写一个函数来计算它们的交集。
// https://leetcode-cn.com/problems/intersection-of-two-arrays/
func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}

// 法一：暴力法，O(m*n)，m和n分别是nums1和nums2的长度

// 法二：使用一个额外的map进行统计，O(n)
func intersection(nums1 []int, nums2 []int) (result []int) {
	m := make(map[int]int, len(nums1))
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if _, exist := m[num]; exist {
			result = append(result, num)
			delete(m, num)
		}
	}
	return result
}
