/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import "fmt"

func main() {
	arr := []int{0,1,2,3,4}
	fmt.Println(arr[3:4]) //只有第四个数,也就是下标3,3
	// --怎么理解 arr[j-1,j],第j个数,下标为j-1
	// --怎么理解 arr[i,j],下标i,-->下标j-1,不包含最后一个j的下标  数量为j-i个

	fmt.Println(arr[len(arr) -1: ])
	fmt.Println(arr[:len(arr) -1])
}

