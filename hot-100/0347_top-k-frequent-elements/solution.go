/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-25 13:48:04
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-25 14:07:05
 * @Description:
 */

package main

import "container/heap"

type Element struct {
	num   int
	count int
}

type PriorityQueue []*Element

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].count > pq[j].count // 构建大顶堆
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{} ) {
	element := x.(*Element)
	*pq = append(*pq, element)
}

func (pq *PriorityQueue) Pop() interface{} {
	element := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return element
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	pq := make(PriorityQueue, 0)
	for num, count := range countMap {
		pq = append(pq, &Element{
			num:   num,
			count: count,
		})
	}

	heap.Init(&pq)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&pq).(*Element).num
	}

	return res
}
