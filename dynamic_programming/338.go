package main

import "fmt"

// 338. 比特位计数
// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
// 进阶:
//     给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
//     要求算法的空间复杂度为O(n)。
//     你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
// https://leetcode-cn.com/problems/counting-bits
func main() {
	// fmt.Println(countBits(5)) // [0,1,1,2,1,2]
	fmt.Println(countBits(8)) // [0,1,1,2,1,2,2,3,1]
}

// 法一：动态规划
// 奇数：dp[i] = dp[i-1]+1
// 偶数：dp[i] = dp[i>>1]
func countBits(num int) []int {
	dp := make([]int, num+1)
	if num == 0 {
		return dp
	}
	dp[1] = 1
	if num == 1 {
		return dp
	}
	for i := 2; i <= num; i++ {
		if (i & 1) == 1 { // 奇数
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i>>1]
		}
	}
	return dp
}

// 法二：二进制
func countBits2(num int) []int {
	dp := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// i & (i-1) 清零最低位的1
		dp[i] = dp[i&(i-1)] + 1
	}

	return dp
}
