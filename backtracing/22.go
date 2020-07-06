package main

import "fmt"

// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// https://leetcode-cn.com/problems/generate-parentheses/
func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis3(3)) // best
}

// 法一：生成全排列，再检查生成字符串是否合法，复杂度较高

// 法二：回溯
// 优化思路：将字符串拼接部分进行优化，选用 []byte 或 []uint8 进行字符串生成
func generateParenthesis(n int) (result []string) {
	if n <= 0 {
		return nil
	}
	generateParenthesisHelper("(", &result, n-1, n)
	return result
}
func generateParenthesisHelper(cur string, result *[]string, left, right int) {
	if left == 0 && right == 0 {
		*result = append(*result, cur)
		return
	}
	if left > 0 {
		generateParenthesisHelper(cur+"(", result, left-1, right)
	}
	if right > left {
		generateParenthesisHelper(cur+")", result, left, right-1)
	}
}

// 法三：动态规划
// n对括号从n-1对变化而来
// dp[i] = "(" + dp[j] + ")" + dp[i- j - 1] , j = 0, 1, ..., i - 1
// 如：dp[2] = {"("+dp[0]+")"+dp[1], "("+dp[1]+")"+dp[0]}
//     dp[3] = {"("+dp[0]+")"+dp[2], "("+dp[1]+")"+dp[1], "("+dp[2]+")"+dp[0]}
// 根据leetcode的测试用例，空间复杂度略低于法二
func generateParenthesis3(n int) []string {
	if n <= 0 {
		return nil
	}
	state := make([][]string, n+1)
	state[0] = []string{""}
	for i := 1; i <= n; i++ {
		cur := []string{}
		for j := 0; j < i; j++ {
			subState1 := state[j]
			subState2 := state[i-j-1]
			for _, str1 := range subState1 {
				for _, str2 := range subState2 {
					cur = append(cur, "("+str1+")"+str2)
				}
			}
		}
		state[i] = cur
	}
	return state[n]
}
