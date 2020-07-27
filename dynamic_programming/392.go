package main

import "fmt"

// 392. 判断子序列
// 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
// 你可以认为 s 和 t 中仅包含英文小写字母。字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
// 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
// 后续挑战 :
// 如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
// https://leetcode-cn.com/problems/is-subsequence/
func main() {
	// fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	// fmt.Println(isSubsequence("axc", "ahbgdc")) // false

	fmt.Println(multiSubsequence("abc", "ahbgdc")) // true
}

// 法一：暴力查找
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	si := 0
	for i := 0; i < len(t) && si < len(s); i++ {
		if s[si] == t[i] {
			si++
		}
	}
	return si == len(s)
}

// 法二：map+二分查找
// map下标为a-z，value为对应字母在t中的索引
// 遍历s，查找比上一个字符下标更大的下标

// 后续挑战：对 t 做预处理
// 这也是原题的动态规划解法
// dp[n][26]表示在t[i]位置，下一个[a-z]的字符位置(不含t[i])
func multiSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	t = " " + t // 为了预处理第一个字符
	dp := make([][]int, len(t))
	for k := range dp {
		dp[k] = make([]int, 26)
	}
	var l byte
	for l = 'a'; l <= 'z'; l++ {
		index := -1
		for j := len(t) - 1; j >= 0; j-- {
			dp[j][l-'a'] = index // init
			if t[j] == l {
				index = j
			}
		}
	}
	si := 0 // 待匹配字符串下标
	for si < len(s) {
		si = dp[si][s[si]-'a']
		if si == -1 {
			return false
		}
	}
	return true
}
