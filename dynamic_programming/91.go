package main

import (
	"fmt"
)

// 91. 解码方法
// 一条包含字母 A-Z 的消息通过以下方式进行了编码：
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// 给定一个只包含数字的非空字符串，请计算解码方法的总数。
// https://leetcode-cn.com/problems/decode-ways/
func main() {
	fmt.Println(numDecodings2("0"))    // 0
	fmt.Println(numDecodings2("0001")) // 0
	fmt.Println(numDecodings2("12"))   // 2
	fmt.Println(numDecodings2("226"))  // 3
	fmt.Println(numDecodings2("27"))   // 1
}

// 不合法的字符的解码方法是0

// 法一：动态规划
// dp[n]表示s的前n个字符有多少种解码方式
// 如果s[i] = 0, dp[i] = dp[i-2]
// 如果s[i-1]是1，dp[i] = dp[i-2] + dp[i-1]
// 如果s[i-1]是2，且s[i]为1-6，dp[i] = dp[i-2] + dp[i-1]
// 其他情况，dp[i] = dp[i-1]
func numDecodings(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1 // 为了方便计算
	dp[1] = 1
	for i := 2; i <= n; i++ {
		if s[i-1] == '0' {
			// 出现'0'有两种情况，一种是10或20，一种是当前的'0'不合法
			if s[i-2] != '1' && s[i-2] != '2' {
				return 0
			}
			dp[i] = dp[i-2]
		} else if (s[i-2] == '1') || (s[i-2] == '2' && s[i-1] >= '1' && s[i-1] <= '6') {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}

// 法二：对法一的空间优化，best
func numDecodings2(s string) (cur int) {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	} else if n == 1 {
		return 1
	}
	p, pp := 1, 1
	for i := 1; i < n; i++ {
		if s[i] == '0' {
			// 出现'0'有两种情况，一种是10或20，一种是当前的'0'不合法
			if s[i-1] != '1' && s[i-1] != '2' {
				return 0
			}
			cur = pp
		} else if (s[i-1] == '1') || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = p + pp
		} else {
			cur = p
		}
		p, pp = cur, p
	}
	return cur
}
