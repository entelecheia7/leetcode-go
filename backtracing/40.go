package main

import (
	"fmt"
	"sort"
)

// 40. 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
//     所有数字（包括目标数）都是正整数。
//     解集不能包含重复的组合。
// https://leetcode-cn.com/problems/combination-sum-ii/
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// 回溯
func combinationSum2(candidates []int, target int) (result [][]int) {
	if len(candidates) == 0 {
		return
	}
	sort.Ints(candidates)
	if candidates[0] > target {
		return
	}
	combinationHelper(candidates, []int{}, 0, target, &result)

	return result
}
func combinationHelper(candidates, cur []int, startIndex, target int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	n := len(cur)
	for i := startIndex; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		cur = append(cur, candidates[i])
		combinationHelper(candidates, cur, i+1, target-candidates[i], result)
		cur = cur[:n]
	}
}
