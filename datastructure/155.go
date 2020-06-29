package main

import (
	"fmt"
)

// 155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
// push(x) —— 将元素 x 推入栈中。
// pop() —— 删除栈顶的元素。
// top() —— 获取栈顶元素。
// getMin() —— 检索栈中的最小元素。
// 提示：
//     pop、top 和 getMin 操作总是在 非空栈 上调用。
// https://leetcode-cn.com/problems/min-stack
func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(9)
	stack.Push(7)
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
}

// 法一：借助一个辅助栈来保存对应元素进出栈时的最小值，和数据的栈同步维护

// 法二：将元素入栈时的最小值同步保存在元素struct
type MinStack struct {
	stack []element
}
type element struct {
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	e := element{
		val: x,
		min: x,
	}
	if len(this.stack) > 0 {
		e.min = getMin(e.min, this.stack[len(this.stack)-1].min)
	}
	this.stack = append(this.stack, e)
}

func (this *MinStack) Pop() {
	// check
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	// check
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].val
}

func (this *MinStack) GetMin() int {
	// check
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].min
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
