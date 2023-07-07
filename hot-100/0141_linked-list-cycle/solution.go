/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-04 10:16:51
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 10:22:12
 * @Description:
 */

package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
