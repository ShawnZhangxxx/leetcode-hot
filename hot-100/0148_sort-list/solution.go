/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-04 15:42:25
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 16:11:44
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
func print(l *ListNode) {
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
func main() {
	head := create(1, 3, 2, -1, 9, 8, 6, 7)
	sortList2(head)
	print(head)
}

func sortList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	l1 := sortList2(head)
	l2 := sortList2(mid)
	return merge2(l1, l2)
}
func merge2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		if l1 == nil {
			return l2
		} else {
			return l1
		}
	}
	head := &ListNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}

	return head.Next
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
