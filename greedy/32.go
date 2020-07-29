package main

import "fmt"

// 32. 最长有效括号
// 给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
// https://leetcode-cn.com/problems/longest-valid-parentheses/
func main() {
	fmt.Println(longestValidParentheses("(()"))      // 2
	fmt.Println(longestValidParentheses(")()())"))   // 4
	fmt.Println(longestValidParentheses("()(()"))    // 2
	fmt.Println(longestValidParentheses("(()(((()")) // 2
	fmt.Println(longestValidParentheses("()(())"))   // 6

}

// 可以暴力求解，列举所有子串，判定是否合法，时间复杂度较高，略

// 法一：计数法
// 对左括号和右括号数量进行统计，当个数相等时，记录数目
// 当右括号数量大于左括号时，重新开始一轮统计
// 为了处理 (() 这种情况，再从右向左统计一遍
// best
func longestValidParentheses(s string) (maxLen int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	left, right := 0, 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
			if left == right {
				maxLen = getMax(maxLen, left+right)
			} else if right > left {
				left, right = 0, 0
			}
		}

	}
	left, right = 0, 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
			if left == right {
				maxLen = getMax(maxLen, left+right)
			} else if left > right {
				left, right = 0, 0
			}
		} else {
			right++
		}
	}

	return maxLen
}

// 法二：栈解法
// 左括号直接入栈
// 遇右括号弹出栈顶元素，计算长度；如果栈为空，将当前位置入栈
// 栈顶要么是左括号，要么是最后一个没有被匹配的右括号的下标
func longestValidParentheses2(s string) (maxLen int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	stack := []int{-1} // 为了计算方便括号的长度
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = getMax(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

// 法三：动态规划
// dp[i] 表示以s[i]结尾的最长有效括号长度
// 如果s[i]为右括号，且s[i-1]为左括号，则dp[i] = dp[i-2]+2
// 如果s[i]为右括号，且s[i-1]也为右括号，需要判断dp[i-1]的前一个位置是否为左括号
func longestValidParentheses3(s string) (maxLen int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' { // 处理()()
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if s[i-1] == ')' { // 处理(())、()(())
				if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
					if i-dp[i-1] >= 2 {
						// 除了在dp[i-1]的基础上判断，还要考虑dp[i]这个序列之前是否有合法的序列
						dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
		}
		maxLen = getMax(maxLen, dp[i])
	}
	return maxLen
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
