/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import (
	"fmt"
	"strconv"
)




func main() {
	// 创建一个切片
	a := "123"
	b := 123
	aInt, _ := strconv.Atoi(a) //字符串转整数
	bString := strconv.Itoa(b) //整数转字符串
	fmt.Println(aInt)
	fmt.Println(bString)
}


