/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-04 15:42:25
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 16:11:44
 * @Description:
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// find middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// cut down
	head1 := head
	head2 := slow.Next
	slow.Next = nil

	// merge
	head1 = sortList(head1)
	head2 = sortList(head2)
	head = merge(head1, head2)
	return head
}

func merge(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}

	if h2 == nil {
		return h1
	}

	dummy := &ListNode{}
	p := dummy
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			p.Next = h1
			h1 = h1.Next
		} else {
			p.Next = h2
			h2 = h2.Next
		}

		p = p.Next
	}

	if h1 != nil {
		p.Next = h1
	}

	if h2 != nil {
		p.Next = h2
	}

	return dummy.Next
}
