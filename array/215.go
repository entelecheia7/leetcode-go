package main

import (
	"container/heap"
	"fmt"
)

// 215. 数组中的第K个最大元素
// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 说明:
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
func main() {
	// fmt.Println((findKthLargest2([]int{3, 2, 1, 5, 6, 4}, 2))) // 5
	fmt.Println((findKthLargest3([]int{3, 2, 1, 5, 6, 4}, 2)))    // 5
	fmt.Println((findKthLargest3([]int{-1, 2, 0}, 3)))            // -1
	fmt.Println((findKthLargest3([]int{-1, 2, 0}, 2)))            // 0
	fmt.Println((findKthLargest3([]int{7, 6, 5, 4, 3, 2, 1}, 5))) // 3
}

// 法一：基于快排，进行优化，O(nlogn)
func findKthLargest(nums []int, k int) int {
	quickSort(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

// 快排，以区间最后一个元素作为基准，结果顺序为递减
func quickSort(nums []int, left, right int, k int) {
	if left >= right {
		return
	}
	pivot := quickSortPartition(nums, left, right)
	if pivot == k-1 {
		return
	} else if pivot > k-1 {
		quickSort(nums, left, pivot-1, k)
	} else {
		quickSort(nums, pivot+1, right, k)
	}
}
func quickSortPartition(nums []int, left, right int) (pivot int) {
	base := nums[right]
	j := left
	// 双指针，将大于base的元素都移动到数组头部
	for i := left; i < right; i++ {
		if nums[i] > base {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
	nums[j], nums[right] = nums[right], nums[j]
	return j
}

// 法二：维护一个容量为k的最小堆维护最大的k个元素
// 当新元素大于堆顶元素才进行添加，O(nlogk)
// 调库函数版本
func findKthLargest2(nums []int, k int) int {
	h := minIntHeap(make([]int, k))
	copy(h, nums[:k])
	heap.Init(&h)
	if len(nums) == k {
		return h.Peek()
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Peek() {
			heap.Pop(&h)
			heap.Push(&h, nums[i])
		}
	}
	return h.Peek()
}

type minIntHeap []int

func (h minIntHeap) Len() int            { return len(h) }
func (h minIntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minIntHeap) Push(x interface{}) { (*h) = append(*h, x.(int)) }
func (h *minIntHeap) Pop() interface{}   { min := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return min }
func (h minIntHeap) Peek() int           { return h[0] }

// 法二的手写最小堆版本
func findKthLargest3(nums []int, k int) int {
	h := minIntHeap(make([]int, k))
	copy(h, nums[:k])
	h.Init()
	if len(nums) == k {
		return h.Peek()
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > h.Peek() {
			h.PopInt()
			h.PushInt(nums[i])
		}
	}
	return h.Peek()
}
func (h minIntHeap) Init() {
	if h.Len() <= 1 {
		return
	}
	for i := (h.Len() >> 1) - 1; i >= 0; i-- {
		h.down(i)
	}
}

func (h *minIntHeap) PopInt() (e int) {
	if h.Len() == 0 {
		return
	}
	e = h.Peek()
	h.Swap(0, h.Len()-1)
	(*h) = (*h)[:h.Len()-1]
	h.down(0)
	return
}

func (h *minIntHeap) PushInt(num int) {
	*h = append(*h, num)
	h.up(h.Len() - 1)
}

// 自下而上堆化
func (h minIntHeap) up(from int) {
	for {
		p := (from - 1) >> 1
		if p < 0 || h.Less(p, from) {
			break
		}
		h.Swap(p, from)
		from = p
	}
}

// 自上而下堆化
func (h minIntHeap) down(from int) {
	for {
		child1 := 2*from + 1
		if child1 >= h.Len() {
			break
		}
		min := -1
		if !h.Less(from, child1) {
			min = child1
		}
		child2 := child1 + 1
		if child2 < h.Len() && h.Less(child2, child1) {
			min = child2
		}
		if min == -1 || h.Less(from, min) {
			break
		}
		h.Swap(min, from)
		from = min
	}
}
