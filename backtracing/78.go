package main

import (
	"fmt"
)

// 78. 子集
// 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
// 说明：解集不能包含重复的子集。
// https://leetcode-cn.com/problems/subsets/
func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets2([]int{1, 2, 3}))
	fmt.Println(subsets3([]int{1, 2, 3}))
}

// 法一：回溯，每个元素不断与它之后的元素组合，形成子集
func subsets(nums []int) (result [][]int) {
	subsetBTHelper(nums, 0, []int{}, &result)
	return result
}
func subsetBTHelper(nums []int, from int, cur []int, result *[][]int) {
	tmp := make([]int, len(cur))
	copy(tmp, cur)
	*result = append(*result, tmp)

	for i := from; i < len(nums); i++ {
		subsetBTHelper(nums, i+1, append(cur, nums[i]), result)
	}
}

// 法二：二进制
// 将 nums 视为长度为 len(nums) 的二进制串 x，则 nums 的子集就是 0-n 位的 x 掩码
// 以 length 为 3 的 nums 为例，其二进制位掩码分别为 000、001、010、011、100、101、111。
// 通过计算其掩码的位是否为 1 (i >> j & 1)来决定是否添加 nums 对于位置的数字到 tmp，得出子集。
func subsets2(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	for i := 0; i < (1 << len(nums)); i++ { // 本层循环用于提供掩码
		tmp := []int{}
		for j := 0; j < len(nums); j++ {
			if (i >> j & 1) == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		result = append(result, tmp)
	}

	return result
}

// 法三：循环，遍历nums时，在结果集的基础上添加元素形成新的子集
func subsets3(nums []int) (result [][]int) {
	if len(nums) == 0 {
		return nil
	}
	result = append(result, []int{})
	for i := 0; i < len(nums); i++ {
		tmp := make([][]int, 0, len(result))
		for _, r := range result {
			tmp = append(tmp, append(r, nums[i]))
		}
		result = append(result, tmp...)
	}
	return result
}
