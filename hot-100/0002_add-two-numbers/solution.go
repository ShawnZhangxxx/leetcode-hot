/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
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
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.



//2<-4<-3
//5<-6<-4
//8<-0<-7
func main()  {
	l1 := create(1,2,3)
	l2 := create(4,5,6)
	l3 := addTwoNumbers2(l1,l2)
	print(l3)
}


func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	carry:= 0
	head := &ListNode{}
	node := head

	for l1!=nil && l2!= nil {
		sum := l1.Val + l2.Val +carry
		carry =  sum / 10
		val := sum % 10   //个位
		newNode := &ListNode{
			Val: val ,
			Next: nil,
		}
		head.Next = newNode
		head = head.Next
		l1	 = l1.Next
		l2  = l2.Next
	}
	if l1 == nil && l2 != nil {
		head.Next = l1
	}
	if l1 != nil && l2 == nil {
		head.Next = l2
	}

	if carry > 0 {
		head.Next = &ListNode{
			Val: carry,
			Next: nil,
		}
	}


	return node.Next
}





func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var (
		head, tail *ListNode
		carry      = 0
	)

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return head
}
