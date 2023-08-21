/*
__author__ = 'robin-luo'
__date__ = '2023/03/17 10:22'
*/

package main

import (
	"fmt"
	"sort"
)

func main()  {
	//[[1,3],[2,6],[8,10],[15,18]]
	nums := [][]int{{1,2},{2,6},{8,10},{15,18}}
	res := merge2(nums)
	fmt.Println(res)
}

func merge2(nums [][]int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})
	var mergedNum = [][]int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i][0] <= mergedNum[len(mergedNum)-1][1] {
			mergedNum[len(mergedNum)-1][1] = nums[i][1]
		}	else {
			mergedNum = append(mergedNum,nums[i])
		}
	}
	return mergedNum
}





type Interval struct {
	Start, End int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	its := make([]Interval, 0)
	for _, it := range intervals {
		its = append(its, Interval{
			Start: it[0],
			End:   it[1],
		})
	}

	sort.Slice(its, func(i, j int) bool {
		return its[i].Start < its[j].Start
	})

	var newIts []Interval
	temp := its[0]
	for i := 1; i < len(its); i++ {
		if its[i].Start <= temp.End {
			temp.End = max(temp.End, its[i].End)
		} else {
			newIts = append(newIts, temp)
			temp = its[i]
		}
	}

	newIts = append(newIts, temp)
	var res [][]int
	for _, newIt := range newIts {
		res = append(res, []int{newIt.Start, newIt.End})
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
