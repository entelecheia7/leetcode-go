package main

import "fmt"

// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。
// 说明：
//     所有数字都是正整数。
//     解集不能包含重复的组合。
// https://leetcode-cn.com/problems/combination-sum-iii/
func main() {
	fmt.Println(combinationSum3(3, 7))
}

func combinationSum3(k int, n int) (result [][]int) {
	combinationHelper([]int{}, k, n, 1, &result)
	return result
}
func combinationHelper(cur []int, k, n, startIndex int, result *[][]int) {
	if k == 0 && n == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	length := len(cur)
	for i := startIndex; i <= 9; i++ {
		if i > n {
			break
		}
		cur = append(cur, i)
		combinationHelper(cur, k-1, n-i, i+1, result)
		cur = cur[:length]
	}
}
