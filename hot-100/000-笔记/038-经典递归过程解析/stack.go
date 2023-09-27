package main

import (
"fmt"
"container/list"
)

func reverseStack(stack *list.List) {
	if stack.Len() == 0 {
		return
	}

	// 弹出栈顶元素
	top := stack.Remove(stack.Back())

	// 递归处理剩余元素
	reverseStack(stack)

	// 将弹出的元素插入到栈底
	stack.PushFront(top)
}

func main() {
	stack := list.New()
	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)
	stack.PushBack(4)
	stack.PushBack(5)

	fmt.Println("原始栈:", stack)

	reverseStack(stack)

	fmt.Println("逆序后的栈:", stack)
}