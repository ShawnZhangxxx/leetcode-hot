/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 14:43'
*/

package main

//2023-08-14
import (
	"container/heap"
	"fmt"
)

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

func main() {
	l1 := create(0, 2, 3)
	l2 := create(4, 5, 6)
	l3 := create(1, 5, 8)
	// l4 := mergeKLists2([]*ListNode{l1, l2, l3})
	// l4 := mergeKLists3([]*ListNode{l1, l2, l3})
	l4 := mergeKLists4([]*ListNode{l1, l2, l3}, 0, 2)
	print(l4)
}

func mergeKLists4(lists []*ListNode, l int, r int) *ListNode { //归并合并
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}

	list1 := mergeKLists4(lists, l, (l+r)/2)
	list2 := mergeKLists4(lists, (l+r)/2+1, r)
	return merge2Lists(list1, list2)

}
func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode { //归并合并
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	head := &ListNode{}
	dummy := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			head = head.Next
			l1 = l1.Next
		} else {
			head.Next = l2
			head = head.Next
			l2 = l2.Next
		}
	}
	if l1 == nil && l2 != nil {
		head.Next = l2
	}

	if l2 == nil && l1 != nil {
		head.Next = l1
	}

	return dummy.Next
}

func mergeKLists3(lists []*ListNode) *ListNode { //不用数组
	temp := minHeap(lists)
	h := &temp
	heap.Init(h)
	dummy := &ListNode{}
	head := dummy
	for h.Len() > 0 {
		p := heap.Pop(h).(*ListNode)
		fmt.Println(p)
		head.Next = p
		head = head.Next
		if p.Next != nil {
			heap.Push(h, p.Next)

		}
		//强转ListNode ,string() 是这么转, 转成 *指针类型
	}
	return dummy.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	temp := minHeap(lists)
	h := &temp
	heap.Init(h)
	dummy := &ListNode{}
	head := dummy
	for h.Len() > 0 {
		p := heap.Pop(h).(*ListNode)
		fmt.Println(p)
		head.Next = p
		head = head.Next
		for _, v := range lists {
			if v == p {
				if v.Next != nil {
					v = v.Next
					heap.Push(h, v)
				}
				break
			}
		}
		//强转ListNode ,string() 是这么转, 转成 *指针类型
	}
	return dummy.Next
}

//要实现heap 需要实现5个方法, 接口必须得用方法来接不是函数 , pop push , sort 里面 的 less ,swap ,len
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

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}

func (m *minHeap) Pop() interface{} {
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
