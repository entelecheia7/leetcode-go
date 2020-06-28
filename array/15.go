package main

import (
	"fmt"
	"sort"
)

// 15. 三数之和
// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// https://leetcode-cn.com/problems/3sum
func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{0, 0, 0}
	fmt.Println(threeSum(nums))
}

// 法一：三层循环，暴力破解

// 法二：双重循环+hash表。是对法一的优化

// 法三：排序+双指针
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)

	result := [][]int{}
	left, right := 0, 0
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right = i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
