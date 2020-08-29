package main

import (
	"fmt"
)

// 214. 最短回文串
// 给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。找到并返回可以用这种方式转换的最短回文串。
// https://leetcode-cn.com/problems/shortest-palindrome/
func main() {
	fmt.Println(shortestPalindrome("aacecaaa")) // aaacecaaa
	fmt.Println(shortestPalindrome("abcd"))     // dcbabcd
	fmt.Println(shortestPalindrome("ba"))       //
}

// 回文串的中心越靠后，需要添加的字符越少，回文串越短，但中心必定在[0:len(s)/2]间
// 法一：暴力，枚举回文串中心，O(n^2)，复杂度差
func shortestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	n := len(s)
	center := (n - 1) / 2
	for i := center; i >= 0; i-- {
		// 以s[i]为奇数长度的中心或偶数长度的右侧中心
		if checkCenter(s, n, i, i) {
			// 字符串头部需要添加的是s[2*i+1:]的倒序字符串，长度是n-2*i-1
			return getReverseStr([]byte(s[2*i+1:])) + s
		}
		if i > 0 && checkCenter(s, n, i-1, i) {
			// 字符串头部需要添加的是s[2*i:]的倒序字符串，长度是n-2*i
			return getReverseStr([]byte(s[2*i:])) + s
		}
	}
	return ""
}
func checkCenter(s string, n int, left, right int) bool {
	for left >= 0 && right < n {
		if s[left] != s[right] {
			return false
		}
		left--
		right++
	}
	return true
}
func getReverseStr(str []byte) string {
	n := len(str)
	if n == 1 {
		return string(str)
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	return string(str)
}

// 法二：KMP todo
// func shortestPalindrome2(s string) string {
// 	n := len(s)
// }
