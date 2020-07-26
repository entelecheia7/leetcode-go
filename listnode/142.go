package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 142. 环形链表 II
// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
// 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
// 说明：不允许修改给定的链表。
// 进阶：
// 你是否可以不用额外空间解决此题？
// https://leetcode-cn.com/problems/linked-list-cycle-ii/
func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	spew.Dump(detectCycle(node1)) // 2
}

// 法一：哈希表保存访问过的节点，O(n)，略

// 法二：快慢指针
// 假设链表有环，当快慢指针相遇
// 令一个指针从head重新开始走，step为1，再次相遇则是环的起点
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	slow = head
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

type ListNode struct {
	Val  int
	Next *ListNode
}
