package main

import "github.com/davecgh/go-spew/spew"

// 83. 删除排序链表中的重复元素
// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func main() {
	spew.Dump(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}))
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
