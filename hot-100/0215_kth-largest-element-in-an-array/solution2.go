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

func main() {
	//heap :=
	res := findKthLargestElement2([]int{1, 5, 3, 4, 2, 6}, 2)
	fmt.Println(res)
}

//这题只能用最xiao堆做
func findKthLargestElement2(nums []int, k int) int {
	ih := &IntHeap{}
	heap.Init(ih)
	for i := 0; i < len(nums); i++ {
		if len(*ih) >= k {
			if nums[i] < (*ih)[0] {
				continue
			} else {
				heap.Push(ih, nums[i])
				heap.Pop(ih)
			}
		} else {
			heap.Push(ih, nums[i])
		}
	}
	fmt.Println(ih)
	return heap.Pop(ih).(int)
}

func (this IntHeap) Len() int {
	return len(this)
}
func (this IntHeap) Less(i int, j int) bool {
	return (this)[i] < (this)[j]
}
func (this IntHeap) Swap(i int, j int) {
	(this)[i], (this)[j] = (this)[j], (this)[i]
}
func (this *IntHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}
func (this *IntHeap) Pop() interface{} {

	pop := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return pop
}
