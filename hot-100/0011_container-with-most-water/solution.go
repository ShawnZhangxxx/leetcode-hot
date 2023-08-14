/*
__author__ = 'robin-luo'
__date__ = '2023/03/08 10:32'
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var res int = maxArea([]int{1, 8, 1, 6, 2, 5, 4, 3, 7})
	fmt.Println(res)
}

func main() {
	maxArea([]int{1,2,3})
}


func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res, area := 0, 0
	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}

		res = int(math.Max(float64(res), float64(area))) //必须转成 float64,之后重新int
	}

	return res
}
