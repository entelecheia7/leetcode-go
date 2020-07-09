package main

import (
	"fmt"
	"math"
)

// 50. Pow(x, n)
// 实现 pow(x, n) ，即计算 x 的 n 次幂函数。
// 说明:
// -100.0 < x < 100.0
// n 是 32 位有符号整数，其数值范围是 [−2^31, 2^31 − 1] 。
// https://leetcode-cn.com/problems/powx-n/
func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow2(2.0, 10))
	// fmt.Println(myPow2(2.0, -2))
	// fmt.Println(myPow(0.00001, 2147483647))

}

// 法一：递归，分治，O(logn)
// 注意考虑-n的溢出情况
func myPow(x float64, n int) float64 {
	if n == 0 || x == 1 {
		return 1
	}
	if n > 0 {
		return myPowHelper(x, n)
	}
	if n == math.MinInt32 {
		x *= x
		n >>= 1
	}
	return 1.0 / myPowHelper(x, -n)
}

func myPowHelper(x float64, n int) (result float64) {
	if n == 0 || x == 1 {
		return 1
	}
	half := myPowHelper(x, n>>1)
	if n%2 == 0 {
		return half * half
	}
	return half * half * x
}

// 法二：循环，二进制解法
// O(logn)
// 图解见：https://leetcode-cn.com/problems/powx-n/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by--15/
func myPow2(x float64, n int) (result float64) {
	if n == 0 || x == 1 {
		return 1
	}
	if n == math.MinInt32 {
		x *= x
		n >>= 1
	}
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	result = 1
	factor := x
	for n > 0 {
		if n%2 == 1 {
			result = result * factor
		}
		factor *= factor
		n >>= 1
	}
	return result
}
