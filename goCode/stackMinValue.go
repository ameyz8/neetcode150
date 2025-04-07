package main

import (
	"fmt"
)

type MinStack struct {
	items []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		items: []int{},
		min:   []int{},
	}
}

func (minStack *MinStack) Push(val int) {
	minStack.items = append(minStack.items, val)
	if len(minStack.min) == 0 || val < minStack.min[len(minStack.min)-1] {
		minStack.min = append(minStack.min, val)
	} else {
		minStack.min = append(minStack.min, minStack.min[len(minStack.min)-1])
	}
}

func (minStack *MinStack) Pop() {
	if len(minStack.items) > 0 {
		minStack.items = minStack.items[:len(minStack.items)-1]
		minStack.min = minStack.min[:len(minStack.min)-1]
	}

}

func (minStack *MinStack) Top() int {
	return minStack.items[len(minStack.items)-1]
}

func (minStack *MinStack) GetMin() int {
	return minStack.min[len(minStack.min)-1]
}

func main() {
	stack := new(MinStack)
	fmt.Println(stack)
	stack.Push(10)
	fmt.Println(stack)
	fmt.Println("Min: ", stack.GetMin())
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Push(3)
	fmt.Println(stack)
	fmt.Println("Min: ", stack.GetMin())
	stack.Push(-10)
	fmt.Println(stack)
	fmt.Println("Min: ", stack.GetMin())
	stack.Push(1)
	fmt.Println(stack)
	fmt.Println("Min: ", stack.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
