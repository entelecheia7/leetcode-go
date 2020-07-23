package main

import "fmt"

// 45. 跳跃游戏 II
// 给定一个非负整数数组，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 说明:
// 假设你总是可以到达数组的最后一个位置。
// https://leetcode-cn.com/problems/jump-game-ii/
func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4})) // 2
}

// 贪心算法
func jump(nums []int) (step int) {
	n := len(nums)
	if n == 1 {
		return 0
	}
	// 因为已经确定终点可达，所以可以直接遍历
	// 只要达到 n-2处，必定可达到终点
	maxPos := 0
	reachable := 0
	for i := 0; i < n-1; i++ {
		maxPos = getMax(maxPos, i+nums[i])
		if i == reachable { // 到达边界时，需要更新步数
			reachable = maxPos
			step++
		}
	}
	return step
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
