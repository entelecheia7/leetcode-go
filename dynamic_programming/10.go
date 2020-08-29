package main

import "fmt"

// 10. 正则表达式匹配
// 给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
// 说明:
//     s 可能为空，且只包含从 a-z 的小写字母。
//     p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
// https://leetcode-cn.com/problems/regular-expression-matching/
func main() {
	// true
	fmt.Println(isMatch2("aa", "aa"))
	fmt.Println(isMatch2("aa", "a."))
	fmt.Println(isMatch2("aa", "a*"))
	fmt.Println(isMatch2("aa", "aa*"))
	fmt.Println(isMatch2("", ".*.*.*"))
	fmt.Println(isMatch2("a", "ab*"))
}

// 法一：分治+递归
// 采用分治思想，对模式串分割成小的单元，再进行匹配
// '.' 可单独为一个单元
// '*'可组合为 .* 、字母*
// 其余字母正常处理
func isMatch(s string, p string) bool {
	if s == "" && p == "" { // 匹配结束
		return true
	} else if (p == "" && s != "") || p[0] == '*' { // * 是模式串首字母，不合法
		// 模式串为空，字符串不为空，一定是无法匹配的
		// 字符串为空，但模式串不为空，是可能匹配成功的，比如 字符+*
		return false
	}
	m, n := len(p), len(s)
	i := m - 1   // 模式串下标
	j := n - 1   // 字符串下标
	for i >= 0 { // 以模式串为基准倒序匹配
		if p[i] == '*' {
			pre := i - 1
			if pre < 0 || p[pre] == '*' { // 不合法的模式串
				return false
			}
			// * 的组合可以匹配字符串s中的0个字符或多个字符，逐一尝试
			var result bool
			if p[i-1] == '.' { //.*
				for j >= -1 { // 支持 s[:0]，即字符串为空的情况
					result = isMatch(s[:j+1], p[:pre])
					if result {
						return true
					}
					j--
				}
			} else { // [a-z] + *
				for j >= 0 && s[j] == p[pre] {
					result = isMatch(s[:j+1], p[:pre])
					if result {
						return true
					}
					j--
				}
				i--
			}
		} else if p[i] == '.' {
			if j < 0 {
				return false
			}
			j--
		} else { // a-z
			if j < 0 || s[j] != p[i] {
				return false
			}
			j--
		}
		i--
	}
	if j >= 0 {
		return false
	}

	return true
}

// 法二：动态规划
// dp[i][j]表示s的前i个字符和p的前j个字符是否匹配
func isMatch2(s string, p string) bool {
	if s == "" && p == "" { // 匹配结束
		return true
	} else if (p == "" && s != "") || p[0] == '*' { // * 是模式串首字母，不合法
		// 模式串为空，字符串不为空，一定是无法匹配的
		// 字符串为空，但模式串不为空，是可能匹配成功的，比如 字符+*
		return false
	}
	m, n := len(s), len(p)
	// dp[i][j] 表示 s[0,i) 是否能匹配 p[0,j)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	// s第i个字符 和 p第j个字符 单个字符匹配函数，不考虑*
	match := func(i, j int) bool {
		if i == 0 && j == 0 {
			return true
		}
		if i == 0 || j == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' { // 字符+*
				dp[i][j] = dp[i][j-2] // 匹配0个
				// 匹配一个
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if match(i, j) { // . 或 字母
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}
