package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 24. 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	// spew.Dump(swapPairs(node))
	spew.Dump(swapPairs1(node))
}

// 法一：循环
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	container := &ListNode{Next: head}
	p := container
	for head != nil && head.Next != nil {
		tail := head.Next.Next
		p.Next = head.Next
		p.Next.Next = head
		head.Next = tail
		p, head = p.Next.Next, head.Next
	}

	return container.Next
}

// 法二：递归
func swapPairs1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head.Next
	head.Next = swapPairs1(head.Next.Next)
	n.Next = head

	return n
}

type ListNode struct {
	Val  int
	Next *ListNode
}
