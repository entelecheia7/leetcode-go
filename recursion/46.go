package main

import (
	"fmt"
)

// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
// https://leetcode-cn.com/problems/permutations/
func main() {
	nums := []int{4, 5, 2, 6}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	result := [][]int{}
	cur := []int{}
	helper(nums, cur, &result)
	return result
}

func helper(nums []int, cur []int, result *[][]int) {
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
		// 排除已经使用的元素
		f := false
		for _, j := range cur {
			if nums[i] == j {
				f = true
				break
			}
		}
		if f {
			continue
		}
		// 添加一个元素
		cur = append(cur, nums[i])
		// 进入下一层
		helper(nums, cur, result)
		// 撤销选择
		cur = cur[:curN]
	}

}
