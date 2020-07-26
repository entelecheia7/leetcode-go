package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 25. K 个一组翻转链表
// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
// 说明：
//     你的算法只能使用常数的额外空间。
//     你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
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

	spew.Dump(reverseKGroup2(node, 2))

	// spew.Dump(reverseKGroup(node, 2))
}

// 法一：递归，先检查节点数，再反转，O(2n)
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 检查是否有k个节点
	i := 1
	tail := head
	for ; i <= k && tail != nil; i++ {
		tail = tail.Next
	}
	if i <= k { // 不足k个节点
		return head
	}
	tailNode := reverseKGroup(tail, k)
	cur := head
	for j := 0; j < k; j++ {
		nextC, nextTail := cur.Next, cur
		cur.Next = tailNode
		cur, tailNode = nextC, nextTail
	}
	return tailNode
}

// 法二：循环，todo

type ListNode struct {
	Val  int
	Next *ListNode
}
