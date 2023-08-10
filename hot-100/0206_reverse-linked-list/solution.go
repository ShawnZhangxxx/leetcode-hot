/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 10:19:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 10:35:09
 * @Description:
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
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
