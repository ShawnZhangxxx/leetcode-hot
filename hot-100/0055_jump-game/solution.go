/*
__author__ = 'robin-luo'
__date__ = '2023/03/15 10:55'
*/

package main

import (
	"fmt"
	"math"
)

func main()  {
	res := canJump2([]int{3,3,1,0,4})
	fmt.Println(res)
}
func canJump2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	lastMax := 0
	for i := 0; i < len(nums) - 1; i++ {
		lastMax = int(math.Max(float64(nums[i]),float64(lastMax - 1)))
	}
	if lastMax > 0 {
		return true
	}
	return false
}


func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	cover := 0
	for i := 0; i <= cover; i++ {
		cover = int(math.Max(float64(cover), float64(i+nums[i])))
		if cover >= len(nums)-1 {
			return true
		}
	}

	return false
}
