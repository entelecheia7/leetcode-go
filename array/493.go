package main

import (
	"fmt"
	"sort"
)

// 493. 翻转对
// 给定一个数组 nums ，如果 i < j 且 nums[i] > 2*nums[j] 我们就将 (i, j) 称作一个重要翻转对。
// 你需要返回给定数组中的重要翻转对的数量。
// 注意:
//     给定数组的长度不会超过50000。
//     输入数组中的所有数字都在32位整数的表示范围内。
// https://leetcode-cn.com/problems/reverse-pairs/
func main() {
	fmt.Println(reversePairs4([]int{2, 4, 3, 5, 1})) // 3
	// fmt.Println(reversePairs3([]int{1, 3, 2, 3, 1})) // 2
}

// 法一：暴力，两两判断，O(n^2)

// 法二：BST
// 问题转换为在元素 nums[i] 的左侧子数组里查找大于 2*nums[i] 的元素
// 时间复杂度高，无法通过大数据测试用例
func reversePairs(nums []int) (count int) {
	var root *Node
	for _, num := range nums {
		count += search(root, 2*num+1)
		root = insert(root, num)
	}
	return count
}
func insert(root *Node, val int) *Node {
	if root == nil {
		root = &Node{Val: val, Count: 1}
	} else if root.Val == val {
		root.Count++
	} else if root.Val < val {
		root.Count++
		root.Right = insert(root.Right, val)
	} else {
		root.Left = insert(root.Left, val)
	}
	return root
}

// 查找root中大于等于val的节点数量
func search(root *Node, val int) int {
	if root == nil {
		return 0
	} else if root.Val < val {
		return search(root.Right, val)
	} else if root.Val > val {
		return root.Count + search(root.Left, val)
	}
	return root.Count
}

type Node struct {
	Val         int // 表示元素的值
	Count       int // 表示当前树中大于等于Val的节点个数（含本节点）
	Left, Right *Node
}

// 法三：BIT树状数组
func reversePairs3(nums []int) (count int) {
	tmp := make([]int, len(nums)) // 保存排序后的 nums 数组
	copy(tmp, nums)
	bit := make([]int, len(nums)+1)
	sort.Ints(tmp)

	for _, num := range nums {
		count += search3(bit, index(tmp, 2*num+1))
		insert3(bit, index(tmp, num))
	}
	return
}
func search3(bit []int, i int) (count int) {
	for i < len(bit) {
		count += bit[i]
		i += i & -i
	}
	return
}
func insert3(bit []int, i int) {
	for i > 0 {
		bit[i] += 1
		i -= i & -i
	}
}

// 如果没找到，则返回比val大的最小元素下标
func index(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] >= val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left + 1
}

// 法四：归并排序
func reversePairs4(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	return reversePairsSub(nums, 0, len(nums)-1)
}
func reversePairsSub(nums []int, left, right int) (count int) {
	if left >= right {
		return 0
	}
	mid := left + ((right - left) >> 1)
	count = reversePairsSub(nums, left, mid) + reversePairsSub(nums, mid+1, right)
	// 计算两个分区的元素对并合两个并排序区间
	merged := make([]int, right-left+1)
	k := 0
	i := left    // i 代表左区间的下标
	j := mid + 1 // j 代表右区间的下标，用于计算元素对
	q := j       // q 代表右区间的下标，用于合并区间
	for i <= mid {
		// 计算nums[i] 和右侧区间的元素对个数
		for j <= right && nums[i] > 2*nums[j] {
			j++
		}
		count += j - (mid + 1)
		// 进行合并
		for q <= right && nums[i] >= nums[q] {
			merged[k] = nums[q]
			k++
			q++
		}
		merged[k] = nums[i]
		k++
		i++
	}
	if q <= right {
		copy(merged[k:], nums[q:])
	}
	copy(nums[left:right+1], merged)
	return count
}
