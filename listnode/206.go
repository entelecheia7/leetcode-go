package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 206. 反转链表
// 反转一个单链表。
// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL
// 进阶:
// 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
// https://leetcode-cn.com/problems/reverse-linked-list/
func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	spew.Dump(reverseList2(node))
}

// 法一：递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// tail 是最后一个节点
	tail := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return tail
}

// 法二：迭代
// best
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		nextP, nextC := cur, cur.Next
		cur.Next = pre
		pre, cur = nextP, nextC
	}
	return pre
}

type ListNode struct {
	Val  int
	Next *ListNode
}
