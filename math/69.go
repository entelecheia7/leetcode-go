package main

import (
	"fmt"
	"math"
)

// 69. x 的平方根
// 实现 int sqrt(int x) 函数。
// 计算并返回 x 的平方根，其中 x 是非负整数。
// 由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
// https://leetcode-cn.com/problems/sqrtx/
func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(6))

}

// 法一：二分法
// O(logn)
func mySqrt(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	left, right := 0, x
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid*mid <= x {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 法二：袖珍计算器算法
// 「袖珍计算器算法」是一种用指数函数 exp⁡\expexp 和对数函数 ln⁡\lnln 代替平方根函数的方法。
// O(1)
func mySqrt2(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	result := int(math.Exp(math.Log(float64(x)) * 0.5))
	if (result+1)*(result+1) <= x {
		return result + 1
	}
	return result
}

// 法三：牛顿迭代法
// 要求根号x的近似值：首先随便猜一个近似值y，然后不断令y等于y和x/y的平均数
// http://www.matrix67.com/blog/archives/361
// O(logn)，二次收敛，因此比二分法快
// best
func mySqrt3(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	y := x >> 1
	for {
		tmp := y * y
		if (tmp == x) || (tmp < x && (y+1)*(y+1) > x) {
			break
		} else {
			y = (y + x/y) >> 1
		}
	}
	return y
}

// 法四：位运算
func mySqrt4(x int) (result int) {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	bit := 1 << 16
	for bit > 0 {
		result |= bit
		if result*result > x {
			result ^= bit
		}
		bit >>= 1
	}
	return result
}
