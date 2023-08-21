/*
__author__ = 'robin-luo'
__date__ = '2023/03/14 17:35'
*/

package main

import (
	"fmt"
	"math"
)

func main()  {
	res := maxSubArray2([]int{1,3,-1,-8,6,7,1})
	fmt.Println(res)
}

func maxSubArray2(nums []int) int {
	res := nums[0]
	var sum int
	for _, v := range nums {
		if sum > 0 {
			sum = sum + v
		}else {
			sum = v
		}
		res = int(math.Max(float64(sum),float64(res)))
	}

	return res
}




func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		res = int(math.Max(float64(res), float64(sum)))
	}

	return res
}
