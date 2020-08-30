package main

import "fmt"

// 115. 不同的子序列
// 给定一个字符串 S 和一个字符串 T，计算在 S 的子序列中 T 出现的个数。
// 一个字符串的一个子序列是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
// 题目数据保证答案符合 32 位带符号整数范围。
// https://leetcode-cn.com/problems/distinct-subsequences/
func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit")) // 3
}

// 法一：动态规划
// dp[i][j]表示t的前i个字符在s的前j个字符的子序列中的个数
// 当s[j-1] != t[i-1]，dp[i][j]直接取t[:i]和s[:j-1]的结果
func numDistinct(s string, t string) int {
	m, n := len(t), len(s)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for k := range dp {
		dp[k] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ { // 空字符串可以视为任何字符串的子序列
		dp[0][j] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[j-1] == t[i-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				// s中的字符是可以删除的，因此dp[i][j]由dp[i][j-1]推出
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}

// 法二：可以对法一进行空间优化，略
