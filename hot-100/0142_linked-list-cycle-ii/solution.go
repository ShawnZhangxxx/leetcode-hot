/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-04 14:58:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 15:04:45
 * @Description:
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}

			return head
		}
	}

	return nil
}
