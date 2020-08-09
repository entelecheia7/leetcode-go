package main

import "fmt"

// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// https://leetcode-cn.com/problems/generate-parentheses/
func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis3(3)) // best
	fmt.Println(generateParenthesis4(3)) // best
}

// 法一：生成全排列，再检查生成字符串是否合法，复杂度较高

// 法二：回溯
// 将字符串拼接部分进行优化，选用 []byte 或 []uint8 进行字符串生成
// best
func generateParenthesis(n int) (result []string) {
	if n <= 0 {
		return nil
	}
	cur := make([]byte, n*2)
	generateHelper(n, n, cur, 0, &result)
	return result
}
func generateHelper(left, right int, cur []byte, i int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, string(cur))
		return
	}
	if left > 0 {
		cur[i] = '('
		generateHelper(left-1, right, cur, i+1, result)
	}
	if right > left {
		cur[i] = ')'
		generateHelper(left, right-1, cur, i+1, result)
	}
}

// 法三：动态规划
// n对括号从n-1对变化而来
// dp[i] = "(" + dp[j] + ")" + dp[i- j - 1] , j = 0, 1, ..., i - 1
// 如：dp[2] = {
// 	   "("+dp[0]+")"+dp[1],
// 	   "("+dp[1]+")"+dp[0]
// }
//     dp[3] = {
// 	   "("+dp[0]+")"+dp[2],
// 	   "("+dp[1]+")"+dp[1],
// 	   "("+dp[2]+")"+dp[0]
// }
// 空间复杂度略低于法二
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

// 法四：广度优先搜索
// 空间复杂度O(n^2)
func generateParenthesis4(n int) (result []string) {
	if n <= 0 {
		return nil
	}
	queue := []generator{{"(", n - 1, n}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.left == 0 && cur.right == 0 {
			result = append(result, cur.cur)
			continue
		}
		if cur.left > 0 {
			queue = append(queue, generator{cur.cur + "(", cur.left - 1, cur.right})
		}
		if cur.right > cur.left {
			queue = append(queue, generator{cur.cur + ")", cur.left, cur.right - 1})
		}
	}

	return result
}

type generator struct {
	cur         string
	left, right int
}
