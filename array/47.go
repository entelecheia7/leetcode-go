package main

import (
	"fmt"
	"sort"
)

// 47. 全排列 II
// 给定一个可包含重复数字的序列，返回所有不重复的全排列。
// https://leetcode-cn.com/problems/permutations-ii/
func main() {
	nums := []int{1, 1, 2}

	fmt.Println(permuteUnique1(nums))
	fmt.Println(permuteUnique2(nums))
}

// 要点：在cur的同一个位置，不能使用重复数字
// 法一：时间最优
// 法二：空间最优

// 法一：通过排序 和 i > 0 && nums[i] == nums[i-1] && used[i-1] 来判断重复数字
// 使用used记录记录使用过的元素，key是元素在nums中的下标
func permuteUnique1(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := [][]int{}
	n := len(nums)
	cur := []int{}
	used := make(map[int]bool, n)
	sort.Ints(nums)

	helper1(nums, cur, used, &result)

	return result
}
func helper1(nums []int, cur []int, used map[int]bool, result *[][]int) {
	curN := len(cur)
	n := len(nums)
	// 结束本次排列
	if curN == n {
		tmp := make([]int, n)
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	for i := 0; i < n; i++ {
		// 排除使用过的元素
		if used[i] {
			continue
		}
		// 排除重复数字：相同的数字，不能在同一个位置出现
		// 剪枝条件分析：https://leetcode-cn.com/problems/permutations-ii/solution/hot-100-47quan-pai-lie-ii-python3-hui-su-kao-lu-zh/
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		//
		cur = append(cur, nums[i])
		used[i] = true
		//
		helper1(nums, cur, used, result)
		//
		cur = cur[:curN]
		used[i] = false
	}
}

// 空间最优
func permuteUnique2(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := [][]int{}
	sort.Ints(nums)
	helper2(nums, 0, &result)

	return result
}

// i表示本次函数需要放置的元素位置
func helper2(nums []int, i int, result *[][]int) {
	n := len(nums)
	if i == n-1 {
		tmp := make([]int, n)
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字
		if k != i && nums[k] == nums[i] {
			continue
		}
		nums[k], nums[i] = nums[i], nums[k]
		helper2(nums, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		nums[i], nums[k] = nums[k], nums[i]
	}
}
