/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import "fmt"

func main() {
	arr := twoSum([]int{1, 11, 3, 9}, 10)
	fmt.Println(arr)
}

func twoSum(nums []int, target int) []int {
	hashmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := hashmap[target-nums[i]]; ok {
			return []int{j, i}
		} else {
			hashmap[nums[i]] = i
		}
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	m := make(map[int]int, len(nums))
// 	for i, num := range nums {
// 		if j, ok := m[target-num]; ok {
// 			return []int{j, i}
// 		}

// 		m[num] = i
// 	}

// 	return nil
// }
