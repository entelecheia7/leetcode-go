package main

import "fmt"

// 44. 通配符匹配
// 给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。
// '?' 可以匹配任何单个字符。
// '*' 可以匹配任意字符串（包括空字符串）。
// 两个字符串完全匹配才算匹配成功。
// 说明:
//     s 可能为空，且只包含从 a-z 的小写字母。
//     p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。
// https://leetcode-cn.com/problems/wildcard-matching/
func main() {
	fmt.Println(isMatch("aa", "a")) // false
	fmt.Println(isMatch("aa", "*")) // true
}

// 法一：动态规划
// dp[i][j]表示s的前i个字符是否可以匹配p的前j个字符
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	// init
	for k := range dp {
		dp[k] = make([]bool, n+1)
	}
	// 当s为空时
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ { // p 不为空，所以从1开始
			if s[i-1] == p[j-1] || p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' { // *可以匹配空字符串或者任意多个字符
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}
