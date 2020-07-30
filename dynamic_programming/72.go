package main

import "fmt"

// 72. 编辑距离
// 给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。
// 你可以对一个单词进行如下三种操作：
//     插入一个字符
//     删除一个字符
//     替换一个字符
// https://leetcode-cn.com/problems/edit-distance/
func main() {
	fmt.Println(minDistance2("horse", "ros")) // 3
}

// 动态规划

// 莱文斯坦距离（Levenshtein distance）
// dp[i][j]代表 word1 的[0:i]到 word2 的[0:j] 的编辑距离
// dp[i][j]总是从 dp[i-1][j]、dp[i][j-1]、dp[i-1][j-1] 变化而来
// dp[i-1][j] => dp[i][j] 需要添加
// dp[i][j-1] => dp[i][j] 需要添加
// dp[i-1][j-1] => dp[i][j] 需要替换
// 如果 word1[i] == word2[j]，dp[i][j] = getMin(dp[i-1][j-1], dp[i-1][j]+1, dp[i][j-1]+1)
// 如果 word1[i] != word2[j]，dp[i][j] = getMin(dp[i-1][j-1]+1, dp[i-1][j]+1, dp[i][j-1]+1)
func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	} else if l2 == 0 {
		return l1
	}
	// init
	// 初始化 word1[0]到word2[0…l2]的编辑距离
	dp := make([][]int, l1)
	for i := 0; i < l1; i++ {
		dp[i] = make([]int, l2)
		if i == 0 {
			for j := 0; j < l2; j++ {
				if word1[0] == word2[j] {
					dp[0][j] = j
				} else if j == 0 {
					dp[0][j] = 1
				} else {
					dp[0][j] = dp[0][j-1] + 1
				}
			}
		}
	}
	// 初始化 word1[0…l1]到word2[0]的编辑距离
	for i := 0; i < l1; i++ {
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else if i > 0 {
			dp[i][0] = dp[i-1][0] + 1
		}
	}

	// 开始递推
	for i := 1; i < l1; i++ {
		for j := 1; j < l2; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin(getMin(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}

	return dp[l1-1][l2-1]
}

// 对以上代码写法进行简化，添加哨兵。时间复杂度没有提升
// dp[i][j]代表 word1 的[0:i)到 word2 的[0:j) 的编辑距离
func minDistance2(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 {
		return l2
	} else if l2 == 0 {
		return l1
	}
	// init
	// 初始化 word1[:0……l1]到word2[:0]的编辑距离
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	// 初始化 word1[:0]到word2[:0……l2]的编辑距离
	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	// 开始递推
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin(getMin(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}

	return dp[l1][l2]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
