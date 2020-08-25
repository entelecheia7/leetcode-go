package main

import "fmt"

// 5. 最长回文子串
// 给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
// https://leetcode-cn.com/problems/longest-palindromic-substring/
func main() {
	fmt.Println(longestPalindrome2("baba"))
}

// 法一：暴力+双指针
// best
func longestPalindrome(s string) (result string) {
	n := len(s)
	if n == 0 {
		return
	}
	result = s[:1]
	for i := 1; i < n-len(result)/2; i++ {
		// 以s[i]为奇数长度回文串的中心
		left, right := i-1, i+1
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if left < i-1 && right-left-1 > len(result) {
			result = s[left+1 : right]
		}
		// 以s[i]为偶数长度回文串的右中心
		left, right = i-1, i
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		if right-left > 1 && right-left-1 > len(result) {
			result = s[left+1 : right]
		}
	}
	return result
}

// 法二：动态规划
// dp[i][j]表示s[i:j]是否为回文字符串，含两侧边界
// O(n^2)
// 该解法还可以进行优化，dp[i][j]仅和sp[i+1][j-1]相关
func longestPalindrome2(s string) (result string) {
	n := len(s)
	if n <= 1 {
		return s
	}
	result = s[:1]
	dp := make([][]bool, n)
	for k := range dp {
		dp[k] = make([]bool, n)
		dp[k][k] = true
	}
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j == i+1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
				if dp[i][j] && j-i+1 > len(result) {
					result = s[i : j+1]
				}
			}
		}
	}
	return result
}
