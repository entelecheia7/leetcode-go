package main

import "fmt"

// 647. 回文子串
// 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
// 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被计为是不同的子串。
// 注意:
//     输入的字符串长度不会超过1000。
// https://leetcode-cn.com/problems/palindromic-substrings/
func main() {
	fmt.Println(countSubstrings("aaa"))    // 6
	fmt.Println(countSubstrings("abc"))    // 3
	fmt.Println(countSubstrings("abac"))   // 5
	fmt.Println(countSubstrings("fdsklf")) // 6
}

// 法一：暴力，best
// 以s中的每个字符为回文串中点，检查计算
func countSubstrings(s string) (count int) {
	n := len(s)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	for i := 0; i < n; i++ {
		count++
		// 回文串长度为奇数
		left, right := i-1, i+1
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}
			count++
			left--
			right++
		}
		// 回文串长度为偶数，以s[i]为中点左侧字符
		left, right = i, i+1
		for left >= 0 && right < n {
			if s[left] != s[right] {
				break
			}
			count++
			left--
			right++
		}
	}
	return count
}

// 法二：动态规划
// 长度更长的回文串总是在长度稍短的回文串的基础上形成
// dp[i][j]表示 s[i:j]是否为回文子串
// dp[i][j]在 dp[i+1][j-1]的基础上判断扩展
func countSubstrings2(s string) (count int) {
	n := len(s)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	dp := make([][]bool, n)
	for k := range dp {
		dp[k] = make([]bool, n)
	}
	for j := 0; j < n; j++ { // j包裹i循环是为了保证 dp[i+1][j-1] 已经算出
		for i := j; i >= 0; i-- {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}
