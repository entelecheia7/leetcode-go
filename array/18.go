package main

import (
	"fmt"
	"sort"
)

// 18. 四数之和
// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。
// https://leetcode-cn.com/problems/4sum/
func main() {
	// fmt.Print(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Print(fourSum([]int{-1, -5, -5, -3, 2, 5, 0, 4}, -7))
}

// 法一：排序，固定两个数+双指针
// 比起三数之和增加了剪枝的步骤
func fourSum(nums []int, target int) (result [][]int) {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		/*剪枝*/
		// 当前范围内的最小四数之和
		min := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if min > target {
			break
		}
		// 当前范围内的最大四数之和
		max := nums[i] + nums[n-1] + nums[n-2] + nums[n-3]
		if max < target {
			continue
		}
		/*剪枝结束*/
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := target - nums[i] - nums[j]
			left, right := j+1, n-1
			for left < right {
				tmp := nums[left] + nums[right]
				if tmp == sum {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if tmp > sum {
					right--
				} else {
					left++
				}
				for left < right && left > j+1 && nums[left] == nums[left-1] {
					left++
				}
				for left < right && right > n-1 && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return
}

// 法二：暴力。复杂度很高
