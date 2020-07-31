package main

import (
	"fmt"
	"math"
)

// 410. 分割数组的最大值
// 给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。
// 注意:
// 数组长度 n 满足以下条件:
//     1 ≤ n ≤ 1000
//     1 ≤ m ≤ min(50, n)
// https://leetcode-cn.com/problems/split-array-largest-sum/
func main() {
	fmt.Println(splitArray2([]int{7, 2, 5, 10, 8}, 2)) // 18
}

// 法一：动态规划，O(n^3)
// dp[i][j]表示以将nums的前i个数分为 j 组得到的最大连续子数组和的最小值(j<=i)
// 设前 k 个数分为 j-1 组，最后一组为第 k+1 到第 i 个数 (k+1>=j)
func splitArray(nums []int, m int) (result int) {
	n := len(nums)
	dp := make([][]int, n+1)
	// init
	for k := range dp {
		dp[k] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[k][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0
	// sums 表示前i个元素的和
	sums := make([]int, n+1)
	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m && j <= i; j++ {
			for k := 0; k < i; k++ {
				dp[i][j] = getMin(dp[i][j], getMax(sums[i]-sums[k], dp[k][j-1]))
			}
		}
	}

	return dp[n][m]
}

// 法二：二分查找
// best
// nums子数组的值在 [max(nums), sum(nums)]中间，使用二分查找不断查找接近值
func splitArray2(nums []int, m int) (result int) {
	n := len(nums)
	left := nums[0]
	right := nums[0]
	for i := 1; i < n; i++ {
		right += nums[i]
		if nums[i] > left {
			left = nums[i]
		}
	}
	for left < right {
		mid := left + ((right - left) >> 1)
		// 从nums[0]开始计算一个 subSum
		// 每找出一个subSum <= mid 的子数组，统计数量+1
		count := 1 // 循环中subSum超过mid才进行统计，因此最后一轮没有在循环中统计，所以赋1
		subSum := 0
		for i := 0; i < n; i++ {
			subSum += nums[i]
			if subSum > mid {
				count++
				subSum = nums[i]
			}
		}
		if count > m { // 分组数量过多，说明选择的subSum过小
			left = mid + 1
		} else { // 分组数量小于等于m，说明选择的subSum可能过大
			right = mid
		}
	}

	return left
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
