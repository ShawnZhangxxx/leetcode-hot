/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-04 10:16:51
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 10:22:12
 * @Description:
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
func create(Data ...int) *ListNode {
	node := &ListNode{}
	head := node
	for i := 0; i < len(Data); i++ {
		newNode := &ListNode{
			Val: Data[i],
		}
		head.Next = newNode
		head = head.Next
	}
	return node.Next
}
func print(l *ListNode)  {
	for l.Next != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Print(l.Val)


}
func main()  {
	l1 := create(1,2,3)
	//l1.Next.Next.Next = l1
	res := hasCycle2(l1)
	fmt.Println(res)
}
func hasCycle2(head *ListNode) bool {
	slow, fast := head, head
	for  fast!= nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
