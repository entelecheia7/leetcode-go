package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func main() {
	node1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	node2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
		},
	}
	spew.Dump(mergeTwoLists3(node1, node2))
}

// 法一：递归
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l2.Next, l1)
	return l2
}

// 法二：循环
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	new := &ListNode{}
	head := new
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			new.Next = l1
			l1 = l1.Next
		} else {
			new.Next = l2
			l2 = l2.Next
		}
		new = new.Next
	}
	if l1 != nil {
		new.Next = l1
	}
	if l2 != nil {
		new.Next = l2
	}

	return head.Next
}

// 法三：循环，不创建新链表，将l2合并到l1
func mergeTwoLists3(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{Next: l1}
	pre := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1 = l1.Next
		} else {
			tmp := l2.Next
			l2.Next = l1
			pre.Next = l2
			l2 = tmp
		}
		pre = pre.Next
	}
	if l1 != nil {
		pre.Next = l1
	}
	if l2 != nil {
		pre.Next = l2
	}

	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
