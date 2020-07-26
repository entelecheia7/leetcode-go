package main

import "fmt"

// 55. 跳跃游戏
// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个位置。
// https://leetcode-cn.com/problems/jump-game/
func main() {
	fmt.Println(canJump2([]int{2, 3, 1, 1, 4}))
	// fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	// fmt.Println(canJump([]int{3, 2, 2, 0, 4}))
}

// 法一：贪心，O(n)
// 从第一个格子逐渐更新可以跳到的最远处
// 直至 到达终点 或 尝试完毕，发现终点不可达
func canJump(nums []int) (result bool) {
	n := len(nums)
	if n == 0 {
		return false
	}
	reachable := 0
	for i := 0; i <= reachable && reachable < n; i++ {
		if i+nums[i] > reachable {
			reachable = i + nums[i]
		}
	}
	return reachable >= n-1
}

// 法二：动态规划
// 从倒数第二个元素开始判断是否可以到达终点，逐步更新终点
// O(n^2)
func canJump2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	}
	state := make([]bool, n)
	state[n-1] = true
	// 2, 3, 1, 1, 4
	for i := n - 2; i >= 0; i-- {
		// j代表在i处可能走的步数
		for j := 1; j <= nums[i] && i+j < n; j++ {
			if state[i+j] == true {
				state[i] = true
				break
			}
		}
	}
	return state[0]
}
