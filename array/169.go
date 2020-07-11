package main

import (
	"fmt"
	"sort"
)

// 169. 多数元素
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// https://leetcode-cn.com/problems/majority-element/description/
func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement4([]int{3, 2, 3}))
}

// 法一：借助一个map统计元素频率，O(n)
func majorityElement(nums []int) int {
	freq := len(nums) / 2
	m := make(map[int]int, freq)
	for _, num := range nums {
		m[num]++
		if m[num] > freq {
			return num
		}
	}
	return -1
}

// 法二：排序，则众数一定在nums[n/2]处。
// 时间：O(nlogn)，空间O(logn)
func majorityElement2(nums []int) int {
	freq := len(nums) / 2
	sort.Ints(nums)
	return nums[freq]
}

// 法三：分治
// 如果 n 是 nums 的众数，那么将 nums 分为两部分，n一定是至少一部分的众数
// 时间复杂度O(nlogn)，空间复杂度O(logn)

// 法四：Boyer-Moore 投票算法
// 我们维护一个候选众数 candidate 和它出现的次数 count。初始时 candidate 可以为任意值，count 为 0；
// 我们遍历数组 nums 中的所有元素，对于每个元素 x，在判断 x 之前，如果 count 的值为 0，我们先将 x 的值赋予 candidate，随后我们判断 x：
//     如果 x 与 candidate 相等，那么计数器 count 的值增加 1；
//     如果 x 与 candidate 不等，那么计数器 count 的值减少 1。
// 在遍历完成后，candidate 即为整个数组的众数。
// 时间复杂度O(n)，空间复杂度O(1)
func majorityElement4(nums []int) int {
	candidate, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count++
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
