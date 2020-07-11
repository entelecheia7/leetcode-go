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
	fmt.Println(combine(4, 2)) // best
	fmt.Println(combine2(4, 2))
}
func combine(n int, k int) [][]int {
	if n < 1 || k < 1 || n < k {
		return nil
	}
	cur := make([]int, k, k)
	result := [][]int{}

	helper1(cur, 0, 1, n, k, &result)

	return result
}

// 有两种写法
// 法一：双百
// ci 是在本次函数中，cur中要添加的元素的位置，避免append操作
func helper1(cur []int, ci, start, n, k int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for i := start; i <= n-k+1; i++ {
		cur[ci] = i
		helper1(cur, ci+1, i+1, n, k-1, result)
	}

}

// 法二
func combine2(n int, k int) [][]int {
	if n < 1 || k < 1 || n < k {
		return nil
	}
	cur := make([]int, 0, k)
	result := [][]int{}
	helper2(cur, 1, n, k, &result)

	return result
}

func helper2(cur []int, num, n, k int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	if num > n {
		return
	}
	cur = append(cur, num)
	//
	helper2(cur, num+1, n, k-1, result)
	//
	cur = cur[:len(cur)-1]

	helper2(cur, num+1, n, k, result)
}
