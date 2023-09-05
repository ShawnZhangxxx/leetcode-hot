/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 10:19:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 10:35:09
 * @Description:
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func create(Data ...int )*ListNode  {
	head := &ListNode{}
	dumy := head
	for i := 0; i < len(Data); i++ {
		node := &ListNode{Val: Data[i]}
		head.Next =  node
		head = head.Next
	}
	return dumy.Next
}

func print(l *ListNode)  {
	for l!= nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func main()  {
	l := create(1,2,3,4,5,6,7,8)
	print(l)
	l = reverseList2(l)
	print(l)
}
func reverseList2(head *ListNode) *ListNode {
		var prev *ListNode
		cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return  prev
}


func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}

	return prev
}
