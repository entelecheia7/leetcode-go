package main

import (
	"fmt"
)

// 641. 设计循环双端队列
// 设计实现双端队列。
// 你的实现需要支持以下操作：
//     MyCircularDeque(k)：构造函数,双端队列的大小为k。
//     insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
//     insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
//     deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
//     deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
//     getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
//     getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
//     isEmpty()：检查双端队列是否为空。
//     isFull()：检查双端队列是否满了。
// 提示：
//     所有值的范围为 [1, 1000]
//     操作次数的范围为 [1, 1000]
//     请不要使用内置的双端队列库。
// https://leetcode-cn.com/problems/design-circular-deque/
func main() {
	// q := (Constructor(3))
	// fmt.Println(q.InsertLast(1))  // true
	// fmt.Println(q.InsertLast(2))  // true
	// fmt.Println(q.InsertFront(3)) // true
	// fmt.Println(q.InsertFront(4)) // false
	// fmt.Println(q.GetRear())      // 2
	// fmt.Println(q.IsFull())       // true
	// fmt.Println(q.DeleteLast())   // true
	// fmt.Println(q.InsertFront(4)) // true
	// fmt.Println(q.GetFront())     // 4

	// q := (Constructor(5))
	// fmt.Println(q.InsertFront(7)) // true
	// fmt.Println(q.InsertLast(0))  // true
	// fmt.Println(q.GetFront())     // 7
	// fmt.Println(q.InsertLast(3))  // true
	// fmt.Println(q.GetFront())     // 7
	// fmt.Println(q.InsertFront(9)) // true
	// fmt.Println(q.GetRear())      // 3
	// fmt.Println(q.GetFront())     // 9
	// fmt.Println(q.GetFront())     // 9
	// fmt.Println(q.DeleteLast())   // true
	// fmt.Println(q.GetRear())      // 0

	q := (Constructor(3))
	fmt.Println(q.InsertFront(8)) // true
	fmt.Println(q.InsertLast(8))  // true
	fmt.Println(q.InsertLast(2))  // true
	fmt.Println(q.GetFront())     // 8
	fmt.Println(q.DeleteLast())   // true
	fmt.Println(q.GetRear())      // 8
	fmt.Println(q.InsertFront(9)) // true
	fmt.Println(q.DeleteFront())  // true
	fmt.Println(q.GetRear())      // 8
	fmt.Println(q.InsertLast(2))  // true
	fmt.Println(q.IsFull())       // true
}

// 循环队列需要两个指针指向头和尾
// 一般头指针指向第一个有效数据的位置，尾指针指向下一个插入位置

// 法一：链表实现
// type MyCircularDeque struct {
// 	// head和tail保留一个哨兵节点
// 	head, tail *TwoWayListNode
// 	length     int
// 	capacity   int
// }
// type TwoWayListNode struct {
// 	Val        int
// 	Next, Prev *TwoWayListNode
// }

// /** Initialize your data structure here. Set the size of the deque to be k. */
// func Constructor(k int) MyCircularDeque {
// 	head := &TwoWayListNode{}
// 	tail := &TwoWayListNode{}
// 	head.Next = tail
// 	tail.Prev = head
// 	return MyCircularDeque{
// 		head:     head,
// 		tail:     tail,
// 		length:   0,
// 		capacity: k,
// 	}
// }

// /** Adds an item at the front of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) InsertFront(value int) bool {
// 	if this.IsFull() {
// 		return false
// 	}
// 	node := this.newNode(value)
// 	node.Next = this.head.Next
// 	node.Next.Prev = node
// 	node.Prev = this.head
// 	this.head.Next = node

// 	this.length++
// 	return true
// }

// /** Adds an item at the rear of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) InsertLast(value int) bool {
// 	if this.IsFull() {
// 		return false
// 	}
// 	node := this.newNode(value)
// 	node.Prev = this.tail.Prev
// 	this.tail.Prev.Next = node
// 	node.Next = this.tail
// 	this.tail.Prev = node

// 	this.length++
// 	return true
// }

// /** Deletes an item from the front of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) DeleteFront() bool {
// 	if this.IsEmpty() {
// 		return false
// 	}
// 	tmp := this.head.Next.Next
// 	this.head.Next.Prev = nil
// 	this.head.Next.Next = nil
// 	this.head.Next = tmp
// 	tmp.Prev = this.head

// 	this.length--
// 	return true
// }

// /** Deletes an item from the rear of Deque. Return true if the operation is successful. */
// func (this *MyCircularDeque) DeleteLast() bool {
// 	if this.IsEmpty() {
// 		return false
// 	}
// 	tmp := this.tail.Prev.Prev
// 	this.tail.Prev.Prev = nil
// 	this.tail.Prev.Next = nil
// 	this.tail.Prev = tmp
// 	tmp.Next = this.tail

// 	this.length--
// 	return true
// }

// /** Get the front item from the deque. */
// func (this *MyCircularDeque) GetFront() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	return this.head.Next.Val
// }

// /** Get the last item from the deque. */
// func (this *MyCircularDeque) GetRear() int {
// 	if this.IsEmpty() {
// 		return -1
// 	}
// 	return this.tail.Prev.Val
// }

// /** Checks whether the circular deque is empty or not. */
// func (this *MyCircularDeque) IsEmpty() bool {
// 	if this.length == 0 && this.capacity != 0 {
// 		return true
// 	}
// 	return false
// }

// /** Checks whether the circular deque is full or not. */
// func (this *MyCircularDeque) IsFull() bool {
// 	if this.length == this.capacity || this.capacity == 0 {
// 		return true
// 	}
// 	return false
// }
// func (this *MyCircularDeque) newNode(val int) *TwoWayListNode {
// 	return &TwoWayListNode{Val: val}
// }

// 法二：数组实现
type MyCircularDeque struct {
	capacity   int
	data       []int
	head, tail int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		capacity: k + 1,
		data:     make([]int, k+1),
		head:     0,
		tail:     0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.head = (this.head - 1 + this.capacity) % this.capacity
	this.data[this.head] = value

	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail = (this.tail - 1 + this.capacity) % this.capacity
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.tail-1+this.capacity)%this.capacity]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return (this.tail+1)%this.capacity == this.head
}
