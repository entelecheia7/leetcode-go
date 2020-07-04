package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 264. 丑数 II
// 编写一个程序，找出第 n 个丑数。
// 丑数就是质因数只包含 2, 3, 5 的正整数。
// 说明:
// 1 是丑数。
// n 不超过1690。
// https://leetcode-cn.com/problems/ugly-number-ii/
func main() {
	fmt.Println(nthUglyNumber2(334))
}

// 法一：动态规划
// 时间复杂度O(n)
func nthUglyNumber(n int) int {
	if n <= 0 {
		return 0
	}
	ugly := make([]int, n)
	ugly[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		u := getMin(getMin(ugly[i2]*2, ugly[i3]*3), ugly[i5]*5)
		ugly[i] = u
		if u == ugly[i2]*2 {
			i2++
		}
		if u == ugly[i3]*3 {
			i3++
		}
		if u == ugly[i5]*5 {
			i5++
		}
	}

	return ugly[n-1]
}
func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 法二：利用堆
// 时间复杂度O(n*logn)
func nthUglyNumber2(n int) int {
	if n <= 0 {
		return 0
	}
	h := minIntHeap([]int{1})
	heap.Init(&h)
	num := 0
	for n > 0 {
		for h[0] == num { // 去重
			heap.Pop(&h)
		}
		min := heap.Pop(&h).(int)
		if min*2 <= math.MaxInt64 {
			heap.Push(&h, min*2)
		}
		if min*3 <= math.MaxInt64 {
			heap.Push(&h, min*3)
		}
		if min*5 <= math.MaxInt64 {
			heap.Push(&h, min*5)
		}
		num = min
		n--
	}

	return num
}

type minIntHeap []int

func (h minIntHeap) Len() int            { return len(h) }
func (h minIntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minIntHeap) Push(x interface{}) { (*h) = append(*h, x.(int)) }
func (h *minIntHeap) Pop() interface{}   { min := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return min }
