///*
// * @Author: lawtech0902 584563542@qq.com
// * @Date: 2023-07-13 13:56:21
// * @LastEditors: lawtech0902 584563542@qq.com
// * @LastEditTime: 2023-07-14 10:42:49
// * @Description:
// */
//
package main
//
//import "container/heap"
//
//func findKthLargest(nums []int, k int) int {
//	ih := &IntHeap{}
//	heap.Init(ih)
//	for i := 0; i < k; i++ {
//		heap.Push(ih, nums[i])
//	}
//
//	for i := k; i < len(nums); i++ {
//		if nums[i] < (*ih)[0] {
//			continue
//		}
//
//		heap.Pop(ih)
//		heap.Push(ih, nums[i])
//	}
//
//	return (*ih)[0]
//}
//
//type IntHeap []int
//
//// Len is the number of elements in the collection.
//func (ih IntHeap) Len() int {
//	return len(ih)
//}
//
//func (ih IntHeap) Less(i int, j int) bool {
//	return ih[i] < ih[j]
//}
//
//func (ih IntHeap) Swap(i, j int) {
//	ih[i], ih[j] = ih[j], ih[i]
//}
//
//func (ih *IntHeap) Push(x any) {
//	*ih = append(*ih, x.(int))
//}
//
//func (ih *IntHeap) Pop() any {
//	x := (*ih)[ih.Len()-1]
//	*ih = (*ih)[:ih.Len()-1]
//	return x
//}
