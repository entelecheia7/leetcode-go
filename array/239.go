package main

import (
	"container/heap"
	"fmt"
)

// 239. 滑动窗口最大值
// 给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。
// 进阶：
// 你能在线性时间复杂度内解决此题吗？
// https://leetcode-cn.com/problems/sliding-window-maximum
func main() {
	fmt.Println(maxSlidingWindow2([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow2([]int{7, 2, 4}, 2))

	fmt.Println(maxSlidingWindow3([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow3([]int{7, 2, 4}, 2))
}

// 法一：暴力法，遍历每个滑动窗口。O(k*n)，略

// 法二：使用堆
// O(logn)
func maxSlidingWindow2(nums []int, k int) (result []int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	var h maxIntHeap
	h = make([][2]int, k)
	for i := 0; i < k; i++ {
		h[i] = [2]int{nums[i], i}
	}
	heap.Init(&h)
	result = append(result, h.Peek().([2]int)[0])
	for i := 1; i <= n-k; i++ {
		for h.Peek() != nil && h.Peek().([2]int)[1] < i {
			heap.Pop(&h)
		}
		heap.Push(&h, [2]int{nums[i+k-1], i + k - 1})
		result = append(result, h.Peek().([2]int)[0])
	}
	return result
}

type maxIntHeap [][2]int // [2]int的第一个元素是数组value，第二个元素是数组下标

func (h maxIntHeap) Len() int            { return len(h) }
func (h maxIntHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h maxIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxIntHeap) Push(x interface{}) { (*h) = append(*h, x.([2]int)) }
func (h *maxIntHeap) Pop() interface{}   { min := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return min }
func (h maxIntHeap) Peek() interface{} {
	if h.Len() == 0 {
		return nil
	}
	return h[0]
}

// 法三： 双端队列。维护一个单调递减队列
// 每个滑动窗口的元素下标都要入队，队的长度不会超过滑动窗口
// O(n+k)
func maxSlidingWindow3(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	deque := []int{0}
	result := make([]int, n-k+1)
	for i := 1; i < n; i++ {
		if len(deque) > 0 && deque[0] == i-k { // 剔除已经滑出窗口范围的数字
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] { // 将栈内小于新窗口元素的元素弹出
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i >= k-1 {
			result[i-k+1] = nums[deque[0]]
		}
	}
	return result
}

// 法四：动态规划
// 思路是将数组按照个数k分为几个组，最后一组可能不满足k个
// 设滑动窗口的左右指针为i、j
// left[j]表示从一个分组范围内从左侧边界向右的最大元素
// right[i]表示从一个分组范围内从右侧边界向左的最大元素
// 滑动窗口的最大值就是两个组范围的最大值
// O(n)
func maxSlidingWindow4(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}
	if k == 1 {
		return nums
	}
	left := make([]int, n)
	left[0] = nums[0]
	right := make([]int, n)
	right[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if i%k == 0 { // 分组的左侧边界
			left[i] = nums[i]
		} else {
			left[i] = getMax(left[i-1], nums[i])
		}
		// right数组要从右往左计算
		j := n - 1 - i
		if (j+1)%k == 0 { // 分组的右侧边界
			right[j] = nums[j]
		} else {
			right[j] = getMax(right[j+1], nums[j])
		}

	}
	result := make([]int, n-k+1)
	for i := 0; i <= n-k; i++ {
		result[i] = getMax(right[i], left[i+k-1])
	}
	return result
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
