package main

import (
	"fmt"
)

// 491. 递增子序列
// 给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
// 说明:
//     给定数组的长度不会超过15。
//     数组中的整数范围是 [-100,100]。
//     给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。
// https://leetcode-cn.com/problems/increasing-subsequences
func main() {
	fmt.Println(findSubsequences2([]int{4, 6, 7, 7})) // [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
	// fmt.Println(findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}))

}

// 回溯
func findSubsequences(nums []int) (result [][]int) {
	n := len(nums)
	findSubsequencesHelper(nums, []int{}, -101, 0, n, &result)
	return result
}

// 剪枝的条件是在 cur 的第 i 位元素不能重复
func findSubsequencesHelper(nums, cur []int, lastElem, i, n int, result *[][]int) {
	if i == n {
		return
	}
	if len(cur) > 1 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
	}
	// 当前层的添加元素不能重复
	used := make(map[int]bool)
	for j := i; j < n; j++ {
		if len(cur) == 0 {
			if !used[nums[j]] {
				used[nums[j]] = true
				findSubsequencesHelper(nums, []int{nums[j]}, nums[j], j+1, n, result)
			}
		} else if nums[j] >= lastElem && !used[nums[j]] {
			cur = append(cur, nums[j])
			used[nums[j]] = true
			findSubsequencesHelper(nums, cur, nums[j], j+1, n, result)
			cur = cur[:len(cur)-1]
		}
	}
}
