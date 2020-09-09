package main

import (
	"container/heap"
	"fmt"
)

// 347. 前 K 个高频元素
// 给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
// 提示：
//     你可以假设给定的 k 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
//     你的算法的时间复杂度必须优于 O(n log n) , n 是数组的大小。
//     题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
// 你可以按任意顺序返回答案。
// https://leetcode-cn.com/problems/top-k-frequent-elements/
func main() {
	fmt.Println(topKFrequent2([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent2([]int{-1, -1}, 1))
	fmt.Println(topKFrequent2([]int{1, 2}, 2))
}

// 法一：利用桶排序。先使用map统计频率，再按照频率分为1-m个桶，最后将桶由高到低输出
// 这种方法适合频率分布范围不大且比较平均的情况
func topKFrequent(nums []int, k int) (result []int) {
	frequency := make(map[int]int, k) // 数字=>频率
	for _, num := range nums {
		frequency[num]++
	}
	freq := make(map[int][]int) // 频率=>数字集合
	maxBucket := 0              // 记录一个最高的频率
	for num, f := range frequency {
		freq[f] = append(freq[f], num)
		if f > maxBucket {
			maxBucket = f
		}
	}
	for i := maxBucket; i >= 1 && k > 0; i-- {
		if num, ok := freq[i]; ok {
			if len(num) <= k {
				result = append(result, num...)
				k -= len(num)
			} else {
				result = append(result, num[:k]...)
				k = 0
			}
		}
	}
	return
}

// 法二：最小堆
// 统计频率后维护一个大小为k的最小堆
func topKFrequent2(nums []int, k int) []int {
	frequency := make(map[int]int, k) // 数字=>频率
	for _, num := range nums {
		frequency[num]++
	}
	// 最小堆中保存数字，使用频率作为比较的依据
	h := newminIntHeap(k)
	for num, freq := range frequency {
		if h.Len() < k {
			heap.Push(h, numWithFreq{num, freq})
		} else if h.Peek().(numWithFreq).freq < freq {
			heap.Pop(h)
			heap.Push(h, numWithFreq{num, freq})
		}
	}

	result := make([]int, 0, k)
	for _, v := range h.data {
		result = append(result, v.num)
	}
	return result
}

type minIntHeap struct {
	data []numWithFreq
}
type numWithFreq struct {
	num  int
	freq int
}

func newminIntHeap(capacity int) *minIntHeap {
	return &minIntHeap{data: make([]numWithFreq, 0, capacity)}
}
func (h minIntHeap) Len() int           { return len(h.data) }
func (h minIntHeap) Less(i, j int) bool { return h.data[i].freq < h.data[j].freq }
func (h minIntHeap) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *minIntHeap) Push(x interface{}) {
	h.data = append(h.data, x.(numWithFreq))
}
func (h *minIntHeap) Pop() interface{} {
	min := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return min
}
func (h minIntHeap) Peek() interface{} {
	if h.Len() == 0 {
		return -1
	}
	return h.data[0]
}
