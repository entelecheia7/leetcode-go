package main

import "fmt"

// 167. 两数之和 II - 输入有序数组
// 给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。
// 函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。
// 说明:
//     返回的下标值（index1 和 index2）不是从零开始的。
//     你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) //{1,2}
}

// 法一：双指针
// O(n)
func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

// 法二：固定一个数，对另一个数进行二分查找
// O(nlogn)
// 略
