package main

import (
	"fmt"
)

// 84. 柱状图中最大的矩形
// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。
// https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
func main() {
	// fmt.Println(largestRectangleArea5([]int{2, 1, 5, 6, 2, 3}))       // 10
	fmt.Println(largestRectangleArea5([]int{6, 7, 5, 2, 4, 5, 9, 3})) // 16
}

// 法一：暴力法。有两种解法。
// 一种是固定高度h，那么向两侧找到比h低的柱子后停下，确定宽度
// 一种是固定宽度，进行枚举
// O(n²)，略

// 法二：单调栈。
// 法二是对法一枚举高度的优化。
// 先通过两次遍历获得左侧和右侧比它低的最近柱子的坐标
// O(n)
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	} else if n == 1 {
		return heights[0]
	}
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1 // -1表示左侧柱子都高于height[i]或左侧无柱子
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			right[i] = n // n表示右侧柱子都高于height[i]或右侧无柱子
		} else {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	max := 0
	for i := 0; i < n; i++ {
		max = getMax(max, (right[i]-left[i]-1)*heights[i])
	}
	return max
}

// 法三：对法二进行改造，求最近最低柱子只遍历一次
// 左边界的位置在i入栈时确定，右边界的位置在i出栈时决定
// O(n)
func largestRectangleArea3(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	} else if n == 1 {
		return heights[0]
	}
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			// 和法二不同，这里right求到的是右侧最近的小于等于当前柱子的柱子
			// 假设若干个连续的柱子高度相同，最右侧的柱子是可以求到准确的右边界的
			// 而左边界始终是精确的
			// 因此最后得到的答案是准确的
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	max := 0
	for i := 0; i < n; i++ {
		max = getMax(max, heights[i]*(right[i]-left[i]-1))
	}
	return max
}

// 法四：和法二类似，不使用栈，两次循环计算出left和right
func largestRectangleArea4(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	} else if n == 1 {
		return heights[0]
	}
	left, right := make([]int, n), make([]int, n)
	left[0] = -1
	right[n-1] = n
	for i := 1; i < n; i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = left[p]
		}
		left[i] = p
	}
	for i := n - 2; i >= 0; i-- {
		p := i + 1
		for p < n && heights[p] >= heights[i] {
			p = right[p]
		}
		right[i] = p
	}

	max := 0
	for i := 0; i < n; i++ {
		max = getMax(max, heights[i]*(right[i]-left[i]-1))
	}
	return max
}

// 法五：使用栈，栈中的元素左边界是它的上一个元素，右边界在遍历中获取
func largestRectangleArea5(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	} else if n == 1 {
		return heights[0]
	}
	heights = append(heights, 0)
	stack := []int{-1} // 递增栈，添加哨兵元素
	max := 0
	for i := 0; i <= n; i++ {
		// 如果遇到一个小于栈顶的元素，说明可以开始计算面积，依次弹出计算
		for stack[len(stack)-1] != -1 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i - stack[len(stack)-1] - 1
			max = getMax(max, w*h)
		}
		stack = append(stack, i)
	}
	heights = heights[:n]
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
