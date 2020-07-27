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

func permute(nums []int) (result [][]int) {
	n := len(nums)
	used := make([]bool, n)
	cur := make([]int, n)
	permuteHelper(nums, n, used, cur, 0, &result)
	return result
}

// index代表本次函数中cur需要填充的位置
func permuteHelper(nums []int, n int, used []bool, cur []int, index int, result *[][]int) {
	// 结束本次排列
	if index == n {
		tmp := make([]int, index)
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < n; i++ {
		// 排除已经使用的元素
		if !used[i] {
			// 添加一个元素
			cur[index] = nums[i]
			used[i] = true
			// 进入下一层
			permuteHelper(nums, n, used, cur, index+1, result)
			// 撤销选择
			used[i] = false
		}
	}

}
