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
//这个题干真是没搞懂,example为什么1节点不是相交节点,
//懂了,这个逼题,相交位置必须还得和头节点距离一样 这个解法是快慢指针的解法

//1,1,1
//1,1 他俩的相交节点为


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
