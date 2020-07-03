package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 108. 将有序数组转换为二叉搜索树
// 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
// 链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
func main() {
	spew.Dump(sortedArrayToBST2([]int{-10, -3, 0, 5, 9}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 法一：递归
// 时间复杂度O(n)
// 空间复杂度O(logn)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) >> 1
	tree := &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}

	return tree
}

// 法二：栈+循环
type MyNode struct {
	parent      *TreeNode
	left, right int // 用于构建子树的数组开始下标和结束下标
	isLeft      bool
}

func newMyNode(parent *TreeNode, left, right int, isLeft bool) *MyNode {
	return &MyNode{parent, left, right, isLeft}
}

func sortedArrayToBST2(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := n >> 1
	root := &TreeNode{}
	stack := []*MyNode{
		newMyNode(root, 0, n-1, true),
	}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last.left <= last.right {
			mid = last.left + ((last.right - last.left) >> 1)
			tree := &TreeNode{Val: nums[mid]}
			if last.isLeft {
				last.parent.Left = tree
			} else {
				last.parent.Right = tree
			}
			stack = append(stack, newMyNode(tree, last.left, mid-1, true))
			stack = append(stack, newMyNode(tree, mid+1, last.right, false))
		}
	}

	return root.Left
}
