/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}

		m[num] = i
	}
	return nil
}
func main() {
	 nums :=	twoSum([]int{1,2,9,7},10)
	 fmt.Println(nums)
}