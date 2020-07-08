package main

import (
	"container/list"
	"fmt"
)

// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
// 提示：
//     1 <= values <= 10000
//     最多会对 appendTail、deleteHead 进行 10000 次调用
// 链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
func main() {
	cq := Constructor()
	cq.AppendTail(3)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())
}

// 一个栈用于维护插入操作，一个栈用于维护删除操作
// 法一：切片实现
type CQueue struct {
	in, out []int
}

func Constructor() CQueue {
	return CQueue{
		in:  make([]int, 0, 10000),
		out: make([]int, 0, 10000),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 && len(this.in) == 0 {
		return -1
	}
	if len(this.out) == 0 {
		for i := len(this.in) - 1; i >= 0; i-- {
			this.out = append(this.out, this.in[i])
			this.in = this.in[:i]
		}
	}

	last := len(this.out) - 1
	v := this.out[last]
	this.out = this.out[:last]
	return v
}

// 法二：链表实现
type CQueue2 struct {
	in, out *list.List
}

func Constructor2() CQueue2 {
	return CQueue2{
		in:  list.New(),
		out: list.New(),
	}
}

func (this *CQueue2) AppendTail(value int) {
	this.in.PushBack(value)
}

func (this *CQueue2) DeleteHead() int {
	if this.out.Len() == 0 {
		for this.in.Len() > 0 {
			this.out.PushBack(this.in.Remove(this.in.Back()))
		}
	}
	if this.out.Len() == 0 {
		return -1
	}
	return this.out.Remove(this.out.Back()).(int)
}
