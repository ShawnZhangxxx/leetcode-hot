/*
__author__ = 'robin-luo'
__date__ = '2023/03/09 11:10'
*/

package main

type ListNode struct {
	Val  int
	Next *ListNode
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
