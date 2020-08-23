package main

import "fmt"

// 680. 验证回文字符串 Ⅱ
// 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
// 注意:
// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
// https://leetcode-cn.com/problems/valid-palindrome-ii/
func main() {
	fmt.Println(validPalindrome2("abca"))
}

// 法一：递归
func validPalindrome(s string) bool {
	return validPalindromeHelper(s, false)
}
func validPalindromeHelper(s string, deleted bool) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			if deleted {
				return false
			}
			// 删除s[left]或s[right]
			return validPalindromeHelper(s[left+1:right+1], true) || validPalindromeHelper(s[left:right], true)
		}
		left++
		right--
	}
	return true
}

// 法二：循环
// best
func validPalindrome2(s string) (result bool) {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] == s[right] {
			left++
			right--
			continue
		}
		// 删除s[left]
		l, r := left+1, right
		result = true
		for l < r {
			if s[l] != s[r] {
				result = false
				break
			}
			l++
			r--
		}
		if result {
			return
		}
		// 删除s[right]
		result = true
		l, r = left, right-1
		for l < r {
			if s[l] != s[r] {
				result = false
				break
			}
			l++
			r--
		}
		return
	}
	return true
}
