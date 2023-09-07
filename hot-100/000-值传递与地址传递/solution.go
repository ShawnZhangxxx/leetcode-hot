/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import "fmt"




func main() {
	// 创建一个切片
	numbers := []int{1, 2, 3, 4, 5}

	// 获取切片的地址
	addr := &numbers

	// 输出切片地址和切片内容
	fmt.Printf("切片的地址：%p\n", addr)
	fmt.Printf("切片的内容：%v\n", *addr)
	fmt.Println(addr)
}


func main2() {
	arr := []int{0,1,2,3,4}
	modifyArr(arr)
	fmt.Println(arr)
	addArrPointer(&arr)
	fmt.Println(arr)

}

func modifyArr(arr []int)  { //内层修改切片外面也会变
	arr[0] = 1
}
func addArr(arr []int)  { //内层append后,外层并没有变化,内层使用的新的地址
	arr = append(arr,1)
	fmt.Println(arr)
}
func addArrPointer(arr *[]int)  { //内层append后,外层并没有变化,内层使用的新的地址
	*arr = append(*arr,1)
	//fmt.Println(arr)
}