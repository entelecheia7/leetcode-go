package main

import (
	"fmt"
)

// 77. 组合
// 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
// 示例:
// 输入: n = 4, k = 2
// 输出:
// [
//   [2,4],
//   [3,4],
//   [2,3],
//   [1,2],
//   [1,3],
//   [1,4],
// ]
// https://leetcode-cn.com/problems/combinations
func main() {
	fmt.Println(combine(4, 2))
}

// 回溯
func combine(n int, k int) (result [][]int) {
	if n < 1 || k < 1 || n < k {
		return nil
	}
	cur := make([]int, k)
	helper(cur, 0, 1, n, k, &result)

	return result
}

// i 是在本次函数中，cur中要添加的元素的位置，避免append操作
func helper(cur []int, i, left, right, k int, result *[][]int) {
	if i == k {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	if left > right || right-left+1 < k-i {
		return
	}
	// 选择当前的数字或不选
	cur[i] = left
	helper(cur, i+1, left+1, right, k, result)
	helper(cur, i, left+1, right, k, result)
}
