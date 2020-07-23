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
	fmt.Println(isMatch("aa", "aa"))
	fmt.Println(isMatch("aa", "a."))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("aa", "aa*"))
	fmt.Println(isMatch("", ".*.*.*"))
	fmt.Println(isMatch("a", "ab*"))
}

// 法一：分治+递归
// 采用分治思想，对模式串分割成小的单元，再进行匹配
// '.' 可单独为一个单元
// '*'可组合为 .* 、字母*
// 其余字母正常处理
// best
func isMatch(s string, p string) bool {
	if s == "" && p == "" { // 匹配结束
		return true
	} else if p == "" && s != "" {
		// 模式串为空，字符串不为空，一定是无法匹配的
		// 字符串为空，但模式串不为空，是可能匹配成功的，比如 字符+*
		return false
	} else if p[0] == '*' { // * 是模式串首字母，不合法
		return false
	}
	lp, ls := len(p), len(s)
	i := lp - 1  // 模式串下标
	sk := ls - 1 // 字符串下标
	for i >= 0 {
		if p[i] == '*' {
			pre := i - 1
			if p[pre] == '*' { // 出现 ** ，说明是一个不合法的模式串
				return false
			}
			// * 的组合可以匹配字符串s中的0个字符或多个字符，逐一尝试
			var result bool
			if p[i-1] == '.' { //.*
				for sk >= -1 { // 支持 s[:0]，即字符串为空的情况
					result = isMatch(s[:sk+1], p[:pre])
					if result {
						return true
					}
					sk--
				}
			} else { // [a-z] + *
				for sk >= 0 && s[sk] == p[pre] {
					result = isMatch(s[:sk+1], p[:pre])
					if result {
						return true
					}
					sk--
				}
				i--
			}
		} else if p[i] == '.' {
			if sk < 0 {
				return false
			}
			sk--
		} else { // a-z
			if sk < 0 || s[sk] != p[i] {
				return false
			}
			sk--
		}
		i--
	}
	if sk >= 0 {
		return false
	}

	return true
}

// 法二：动态规划
func isMatch2(s string, p string) bool {
	if s == "" && p == "" {
		return true
	} else if p == "" && s != "" {
		// 模式串为空，字符串不为空，一定是无法匹配的
		// 字符串为空，但模式串不为空，是可能匹配成功的，比如 字符+*
		return false
	} else if p[0] == '*' { // * 是模式串首字母，不合法
		return false
	}
	lp, ls := len(p), len(s)
	// dp[i][j] 表示 s[0,i) 是否能匹配 p[0,j)
	dp := make([][]bool, ls+1)
	for i := 0; i <= ls; i++ {
		dp[i] = make([]bool, lp+1)
	}
	dp[0][0] = true
	// s[sk-1,sk) 和 p[pk-1:pk) 单个字符匹配函数，不考虑*
	match := func(sk, pk int) bool {
		if sk == 0 && pk == 0 {
			return true
		}
		if sk == 0 || pk == 0 {
			return false
		}
		if p[pk-1] == '.' {
			return true
		}
		return s[sk-1] == p[pk-1]
	}

	for sk := 0; sk <= ls; sk++ {
		for pk := 1; pk <= lp; pk++ {
			if p[pk-1] == '*' { // 字符+*
				dp[sk][pk] = dp[sk][pk-2] || dp[sk][pk] // 匹配0个
				// 匹配一个
				if match(sk, pk-1) {
					dp[sk][pk] = dp[sk][pk] || dp[sk-1][pk]
				}
			} else if match(sk, pk) { // . 或 字母
				dp[sk][pk] = dp[sk][pk] || dp[sk-1][pk-1]
			}
		}
	}

	return dp[ls][lp]
}
