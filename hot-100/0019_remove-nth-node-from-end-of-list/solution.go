/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 11:10'
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := create(1, 2, 3, 4)
	l = removeNthFromEnd2(l, 1)
	print(l)
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
func print(l *ListNode) {
	for l.Next != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
	fmt.Print(l.Val)
}
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var f, s *ListNode = head, head
	for i := 0; i < n; i++ {
		f = f.Next
	}
	for f.Next != nil {
		f = f.Next
		s = s.Next
	}
	s.Next = s.Next.Next
	return head

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := head
	slow := dummy
	i := 1
	for fast != nil {
		fast = fast.Next
		if i > n {
			slow = slow.Next
		}

		i++
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}
