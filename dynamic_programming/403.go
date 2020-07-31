package main

import "fmt"

// 403. 青蛙过河
// 一只青蛙想要过河。 假定河流被等分为 x 个单元格，并且在每一个单元格内都有可能放有一石子（也有可能没有）。 青蛙可以跳上石头，但是不可以跳入水中。
// 给定石子的位置列表（用单元格序号升序表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一个石子上）。 开始时， 青蛙默认已站在第一个石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格1跳至单元格2）。
// 如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。
// 请注意：
//     石子的数量 ≥ 2 且 < 1100；
//     每一个石子的位置序号都是一个非负整数，且其 < 231；
//     第一个石子的位置永远是0。
// https://leetcode-cn.com/problems/frog-jump/
func main() {
	// 	跳1个单位到第2块石子, 然后跳2个单位到第3块石子, 接着
	// 跳2个单位到第4块石子, 然后跳3个单位到第6块石子,
	// 跳4个单位到第7块石子, 最后，跳5个单位到第8个石子（即最后一块石子）。
	fmt.Println(canCross([]int{0, 1, 3, 5, 6, 8, 12, 17})) // true
	fmt.Println(canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))  // false
	fmt.Println(canCross([]int{0, 2}))                     // false
}

// 法一：动态规划
// dp[i][j]表示在从某个位置 x 跳 j 步 是否可达stones[i]（1<=j<=i）
// 则到达位置 x 的步数为 j-1 || j || j+1
// dp[i][j] = dp[x][k-1] || dp[x][j] || dp[x][j+1]
// O(n^2)
func canCross(stones []int) bool {
	n := len(stones)
	if n == 2 {
		if stones[1] != 1 {
			return false
		}
		return true
	}
	dp := make([][]bool, n)
	for k := range dp {
		dp[k] = make([]bool, n+1) // +1是为了保证 j+1 步不越界
	}
	dp[1][1] = true
	for i := 2; i < n; i++ {
		// 本次循环求解 dp[i]，stones[i]可以从 1~i-1 位置抵达，列举所有可能
		for x := 1; x < i; x++ {
			// 从 stones[x] 跳到 stones[i]需要多少步
			needStep := stones[i] - stones[x]
			// 如果从 stones[x] 跳 needStep 步可达 stones[i]，则有 dp[x][needStep] 或 dp[x][needStep-1] 或 dp[x][needStep+1] 为 true
			if needStep <= i {
				dp[i][needStep] = dp[x][needStep] || dp[x][needStep-1] || dp[x][needStep+1]
				if i == n-1 && dp[i][needStep] {
					return true
				}
			}
		}

	}

	return false
}

// 法二：DFS记忆化搜索+二分，todo
// O(n^2logn)
