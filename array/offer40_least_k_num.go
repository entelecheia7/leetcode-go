package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 剑指 Offer 40. 最小的k个数
// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
// https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/
func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers2([]int{3, 2, 1}, 2))
	fmt.Println(getLeastNumbers3([]int{3, 2, 1}, 2))
}

// 法一：排序，时间O(nlogn)，空间O(logn)
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

// 法二：利用大顶堆，时间O(nlogk)，空间O(k)
func getLeastNumbers2(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	}
	var h maxIntHeap = make([]int, k)
	copy(h, arr)
	heap.Init(&h)
	for i := k; i < len(arr); i++ {
		if h[0] > arr[i] {
			heap.Pop(&h)
			heap.Push(&h, arr[i])
		}
	}

	return h
}

type maxIntHeap []int

func (h maxIntHeap) Len() int            { return len(h) }
func (h maxIntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxIntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxIntHeap) Push(x interface{}) { (*h) = append(*h, x.(int)) }
func (h *maxIntHeap) Pop() interface{}   { min := (*h)[len(*h)-1]; (*h) = (*h)[:len(*h)-1]; return min }

// 法三：利用快排思想进行递归分区，时间O(n)
// arr[k-1]是期待的分区中点
func getLeastNumbers3(arr []int, k int) []int {
	if len(arr) < k {
		return arr
	} else if k == 0 {
		return nil
	}
	return quickSearch(arr, 0, len(arr)-1, k)
}
func quickSearch(arr []int, left, right, k int) []int {
	mid := partition(arr, left, right)
	if mid == k-1 {
		return arr[:k]
	} else if mid < k-1 {
		return quickSearch(arr, mid+1, right, k)
	}
	return quickSearch(arr, left, mid-1, k)
}

// 返回分区下标
func partition(arr []int, left, right int) (pivot int) {
	standard := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < standard {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[right], arr[i] = arr[i], arr[right]
	return i
}
