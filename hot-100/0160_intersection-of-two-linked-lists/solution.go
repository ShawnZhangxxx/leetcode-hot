/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-06 10:49:19
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-06 10:51:29
 * @Description:
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
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
