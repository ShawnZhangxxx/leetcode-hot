/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 14:43'
*/

package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	if m[i] == nil {
		return false
	}

	if m[j] == nil {
		return true
	}

	return m[i].Val < m[j].Val
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x any) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() any {
	size := len(*m)
	res := (*m)[size-1]
	*m = (*m)[:size-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	temp := minHeap(lists)
	h := &temp
	heap.Init(h)
	var head, last *ListNode
	for h.Len() > 0 {
		v := heap.Pop(h).(*ListNode)
		if v == nil {
			continue
		}

		if last != nil {
			last.Next = v
		}

		if head == nil {
			head = v
		}

		last = v
		if v.Next != nil {
			heap.Push(h, v.Next)
		}
	}

	return head
}
