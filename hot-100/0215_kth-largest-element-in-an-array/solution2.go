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
import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func main()  {
	//heap :=
	res :=	findKthLargestElement([]int{1,5,3,4,2,6},3)
	fmt.Println(res)
}

func findKthLargestElement(nums []int,k int ) int {
	ih := &IntHeap{}
	heap.Init(ih)
	for i := 0; i < len(nums); i++ {
		heap.Push(ih,nums[i])
		if len(*ih) == k+1 {
			heap.Pop(ih)
		}
	}
	return heap.Pop(ih).(int)
}

func  (this IntHeap) Len() int {
	return len(this)
}
func  (this IntHeap) Less(i int,j int) bool  {
	return (this)[i] <  (this)[j]
}
func  (this IntHeap) Swap(i int,j int)  {
	(this)[i] , (this)[j] = (this)[j] , (this)[i]
}
func  (this *IntHeap) Push(x interface{})  {
	*this = append(*this,x.(int))
}
func  (this *IntHeap) Pop() interface{}  {

	pop := (*this)[len(*this) -1]
	*this = (*this)[:len(*this) -1]
	return pop
}