/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 14:25:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-03 14:29:48
 * @Description:
 */

package main

import "fmt"

func main()  {
	fmt.Println(0^0)
	//^ 相同为0不同为1, 0和任何数异或还是那个数,相同为0不同为1
	res := singleNumber([]int{2,2,3,5,3})
	fmt.Println(res)
}

func singleNumber(nums []int) int {
	a := 0
	for _, num := range nums {
		a ^= num
	}

	return a
}
