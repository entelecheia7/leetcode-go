package main

import (
	"github.com/davecgh/go-spew/spew"
)

// 141. 环形链表
// 给定一个链表，判断链表中是否有环。
// https://leetcode-cn.com/problems/linked-list-cycle/
func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	spew.Dump(hasCycle2(node1))
}

// 法一：使用哈希表，记录访问过的指针，遍历时进行判断，空间复杂度高

// 法二：使用快慢指针，如果链表中有环，那么快慢指针一定会走到同一个位置
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}

// 法二的一种行为艺术写法
func hasCycle2(head *ListNode) (result bool) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	fast, slow := head.Next, head
	for fast != slow {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
