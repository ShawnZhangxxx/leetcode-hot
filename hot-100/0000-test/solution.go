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
	array(arr)
	fmt.Println(arr)
}

func array(arr []int)  {
	arr[0] = 1
}