package main

import "fmt"

// 1143. 最长公共子序列
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
// 若这两个字符串没有公共子序列，则返回 0。
// 提示:
//     1 <= text1.length <= 1000
//     1 <= text2.length <= 1000
//     输入的字符串只含有小写英文字符。
// https://leetcode-cn.com/problems/longest-common-subsequence
func main() {
	fmt.Println(longestCommonSubsequence("abc", "abc"))   // 3
	fmt.Println(longestCommonSubsequence("abc", "def"))   // 0
	fmt.Println(longestCommonSubsequence("abcde", "ace")) // 3
}

// 动态规划
// dp[i][j]表示 text1的前i个字符 和 text2的前j个字符 的最长子序列长度
// 求解 dp[i][j]，需要在 dp[i-1][j-1]、dp[i-1][j]、dp[i][j-1]的基础上判断 text1[i]和 text2[j]的关系
// 如果 text1[i-1] == text2[j-1]，则dp[i][j] = dp[i-1][j-1]+1
// 如果 text1[i-1] != text2[j-1]，则dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// 不等的情况下，不考虑dp[i-1][j-1]，这样会漏掉答案，比如abcx和abxc
func longestCommonSubsequence(text1 string, text2 string) int {
	if text1 == "" || text2 == "" {
		return 0
	}
	if text1 == text2 {
		return len(text1)
	}
	l1, l2 := len(text1), len(text2)
	dp := make([][]int, l1+1)
	for k := range dp {
		dp[k] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = getMax(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[l1][l2]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
