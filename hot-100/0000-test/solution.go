/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import "fmt"

func main() {
	arr := []int{2, 11, 3, 9}
	fmt.Println(arr[len(arr) -1: ])
	fmt.Println(arr[:len(arr) -1])
	fmt.Println(arr[:1])
	fmt.Println(arr[1:])
	array1(arr)
	array2(arr)
	fmt.Println(arr)
	arr2 := arr[:]//切片
	arr[0] = 10
	fmt.Println(arr2)
}

func array1(arr []int)  { //内层修改切片外面也会变
	arr[0] = 1
}
func array2(arr []int)  { //内层append后,外层并没有变化,内层使用的新的地址
	arr = append(arr,1)
}