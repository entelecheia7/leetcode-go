package main

import (
	"fmt"
)

// 11. 盛最多水的容器
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
// https://leetcode-cn.com/problems/container-with-most-water
func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}

// 法一：暴力破解
// 计算出所有容器的容量，进行比较。时间复杂度高

// 法二：双指针
// 双指针需要移动的是更短的一侧
func maxArea(height []int) (maxArea int) {
	for left, right := 0, len(height)-1; left < right; {
		maxArea = getMax(maxArea, (right-left)*getMin(height[left], height[right]))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
