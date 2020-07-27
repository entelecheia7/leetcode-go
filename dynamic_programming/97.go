package main

import "fmt"

// 97. 交错字符串
// 给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。
// https://leetcode-cn.com/problems/interleaving-string/
func main() {
	fmt.Println(isInterleave2("aabcc", "dbbca", "aadbbcbcac"))                                                                                                                                                                                                       // true
	fmt.Println(isInterleave2("aabcc", "dbbca", "aadbbbaccc"))                                                                                                                                                                                                       // false
	fmt.Println(isInterleave2("accbaabaaabbcbaacbababacaababbcbabaababcaabbbbbcacbaa", "cabaabcbabcbaaaacababccbbccaaabaacbbaaabccacabaaccbbcbcb", "accbcaaabbaabaaabbcbcbabacbacbababaacaaaaacbabaabbcbccbbabbccaaaaabaabcabbcaabaaabbcbcbbbcacabaaacccbbcbbaacb")) // true
}

// 法一：回溯，时间复杂度较高
func isInterleave(s1 string, s2 string, s3 string) (result bool) {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	isInterleaveHelper(s1, s2, s3, 0, 0, 0, &result)
	return result
}
func isInterleaveHelper(s1, s2, s3 string, k, i, j int, result *bool) bool {
	if *result {
		return true
	} else if k == len(s3) {
		*result = true
		return true
	}
	var r1, r2 bool
	if i < len(s1) {
		if s3[k] != s1[i] {
			r1 = false
		} else {
			r1 = isInterleaveHelper(s1, s2, s3, k+1, i+1, j, result)
		}
	}
	if j < len(s2) {
		if s3[k] != s2[j] && !r1 {
			return false
		}
		r2 = isInterleaveHelper(s1, s2, s3, k+1, i, j+1, result)
	}
	return r1 || r2
}

// 法二：动态规划，best
// dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符是否可以组成 s3[:i+j+1]
func isInterleave2(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l1+l2 != l3 {
		return false
	}
	dp := make([][]bool, l1+1)
	for k := range dp {
		dp[k] = make([]bool, l2+1)
	}
	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && s3[i+j-1] == s2[j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && s3[i-1] == s1[i-1]
			} else {
				dp[i][j] = (dp[i][j-1] && (s3[i+j-1] == s2[j-1]))
				if !dp[i][j] {
					dp[i][j] = dp[i-1][j] && s3[i+j-1] == s1[i-1]

				}
			}
		}
	}

	return dp[l1][l2]
}
