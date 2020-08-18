package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 109. 有序链表转换二叉搜索树
// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
// https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
func main() {
	spew.Dump(sortedListToBST2(&ListNode{
		Val: -10,
		Next: &ListNode{
			Val: -3,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
		},
	}))
}

// 法一：常规思路。找到链表的中点，作为根。递归计算
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
		}
	}
	// 找中点
	copier := &ListNode{Next: head}
	fast, slow := head, copier
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	root := &TreeNode{
		Val:   slow.Next.Val,
		Right: sortedListToBST(slow.Next.Next),
	}
	slow.Next = nil
	root.Left = sortedListToBST(head)
	return root
}

// 法二：写法优化。不对原始链表做改动
func sortedListToBST2(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return sortedListToBSTHelper(head, nil)
}
func sortedListToBSTHelper(head, tail *ListNode) *TreeNode {
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		fast = fast.Next.Next
		slow = slow.Next
	}
	root := &TreeNode{
		Val:   slow.Val,
		Left:  sortedListToBSTHelper(head, slow),
		Right: sortedListToBSTHelper(slow.Next, tail),
	}
	return root
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
