package main

import (
	"fmt"
	"sort"
)

// 39. 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
//     所有数字（包括 target）都是正整数。
//     解集不能包含重复的组合。
// 提示：
//     1 <= candidates.length <= 30
//     1 <= candidates[i] <= 200
//     candidate 中的每个元素都是独一无二的。
//     1 <= target <= 500
// https://leetcode-cn.com/problems/combination-sum
func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[7], [2,2,3]]
}

// 法一：排序+回溯
// 去重的条件是当前的可选元素必须大于等于回溯列表中的最后一个元素
func combinationSum(candidates []int, target int) (result [][]int) {
	sort.Ints(candidates)
	if candidates[0] > target {
		return nil
	}
	combinationSumHelper(candidates, []int{}, 0, target, &result)
	return result
}
func combinationSumHelper(candidates []int, cur []int, candidateIndex int, target int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}
	for i := candidateIndex; i < len(candidates); i++ {
		if candidates[i] > target {
			return
		}
		cur = append(cur, candidates[i])
		combinationSumHelper(candidates, cur, i, target-candidates[i], result)
		cur = cur[:len(cur)-1]
	}
}
