package main

import (
	"fmt"
	"math"
)

// 42. 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// https://leetcode-cn.com/problems/trapping-rain-water/
func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1} //6
	// nums := []int{4, 2, 3} // 1
	fmt.Println(trap(nums))
}

// 法一：按列求。
// 对于每一个列，如果它左侧和右侧最高的墙比它高，这个列就有雨水
// O(n)
func trap(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	// 先计算出对于每一列，它左侧和右侧最高的墙，避免在循环中求解
	maxLeft := make([]int, n)
	maxLeft[0] = math.MinInt64
	for i := 1; i < n; i++ {
		maxLeft[i] = getMax(maxLeft[i-1], height[i-1])
	}
	maxRight := make([]int, n)
	maxRight[n-1] = math.MinInt64
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = getMax(maxRight[i+1], height[i+1])
	}

	capacity := 0
	for i := 1; i < n-1; i++ {
		min := getMin(maxRight[i], maxLeft[i])
		if min > height[i] {
			capacity += min - height[i]
		}
	}

	return capacity
}

// 法二：对法一的优化，maxLeft 和 maxRight 不使用数组
// 使用双指针，从两端计算
// 当maxLeft小于maxRight时，从左侧计算
// 当maxRight小于maxLeft时，从右侧计算
// 从左计算，mexLeft是准确的。对于一个列来说，如果它的maxLeft < height[n-1]，那么它的maxLeft和maxRight的最小值一定在左侧
// 从右计算，maxRight是准确的。对于一个列来说，如果它的maxRight < height[0]，那么它的maxLeft和maxRight的最小值一定在右侧
func trap2(height []int) int {
	n := len(height)
	if n < 3 {
		return 0
	}
	maxLeft := height[0]    // left指针左侧最大元素
	maxRight := height[n-1] // right指针右侧最大元素
	capacity := 0
	left, right := 1, n-2
	for left <= right {
		if maxLeft <= maxRight {
			min := maxLeft
			if min > height[left] {
				capacity += min - height[left]
			}
			maxLeft = getMax(maxLeft, height[left])
			left++
		} else {
			min := maxRight
			if min > height[right] {
				capacity += min - height[right]
			}
			maxRight = getMax(maxRight, height[right])
			right--
		}
	}

	return capacity
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
