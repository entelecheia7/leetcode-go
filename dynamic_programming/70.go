package main

import (
	"fmt"
)

// 70. 爬楼梯
// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
// 注意：给定 n 是一个正整数。
// 链接：https://leetcode-cn.com/problems/climbing-stairs
func main() {
	fmt.Println(climbStairs(3))
}

// 递推公式：f(n) = f(n-1)+f(n-2)
// 实现方式：递归、循环
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	pp := 1 // f(n-2)
	p := 2  // f(n-1)
	for i := 3; i <= n; i++ {
		p, pp = p+pp, p
	}
	return p
}
