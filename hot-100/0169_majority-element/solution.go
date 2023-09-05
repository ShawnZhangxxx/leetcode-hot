/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-07 15:10:35
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-07 16:16:57
 * @Description:
 */

package main

import "fmt"

func main()  {
	res := majorityElement2([]int{1,2,3,4,6,6,1,1,1,1,1,1,6,1})
	fmt.Println(res)
}

func majorityElement2(nums []int)int  {
	major,count := 0,0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			major = nums[i]
		}
		if major == nums[i] {
			count ++
		}else {
			count --
		}

	}
	return  major
}















func majorityElement(nums []int) int {
	major, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			major = num
		}

		if major == num {
			count++
		} else {
			count--
		}
	}

	return major
}
