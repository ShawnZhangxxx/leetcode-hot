/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-18 13:50:26
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-18 14:01:36
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
	l := create(1,2,3,3,1)
	res := isPalindrome(l)
	fmt.Println(res)
	print(l)
}
//多线程条件下要加锁,改变了指针的指向 slow之后所有指针专项
func isPalindrome(head *ListNode) bool {
	headTmp := head
	fast,slow := head,head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tail := reverse(slow)
	tailTmp := tail
	for tail != nil {
		if head.Val != tail.Val {
			//转回去
			reverse(tailTmp)
			return  false
		}
		head = head.Next
		tail = tail.Next
	}
	print(headTmp)

	//转回去
	reverse(tailTmp)
	return true



}
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev


}