/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-06 10:49:19
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-06 10:51:29
 * @Description:
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
//这个题干真是没搞懂,example为什么1节点不是相交节点,
//懂了,这个逼题,相交位置必须还得和头节点距离一样 这个解法是快慢指针的解法

//1,1,1
//1,1 他俩的相交节点为
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
	fmt.Println(l.Val)


}
func main()  {
	////Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2
	////, skipB = 3
	l1 := create(4)
	var i1 = &ListNode{Val: 1,Next: nil}
	var i5 = &ListNode{Val: 7,Next: nil}
	l1.Next = i1
	l1.Next.Next = i5
	print(l1)
	l2 := create(5,6)
	l2.Next.Next = i1
	l2.Next.Next.Next = i5
	print(l2)

	res := getIntersectionNode(l1,l2)
	fmt.Println(res.Val)
}

func getIntersectionNode(h1, h2 *ListNode) *ListNode {
	l1, l2 := h1, h2
	for l1 != l2 {
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = h2
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = h1
		}
	}

	return l1
}
