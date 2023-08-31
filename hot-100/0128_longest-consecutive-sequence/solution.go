/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 13:56:06
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-03 14:24:36
 * @Description:
 */

package main

import (
	"fmt"
	"math"
	"sort"
)

func main()  {
	res:=	longestConsecutive2([]int{1,3,9,4,101,2,5,5,8,6})
	fmt.Println(res)
}


func longestConsecutive2(nums []int) int{
	uf := newUf(len(nums))
	keyMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _,ok := keyMap[nums[i]] ; ok {
			continue
		}
		if _,ok := keyMap[nums[i] - 1] ; ok {
			uf.union(i,keyMap[nums[i] - 1])
		}
		if _,ok := keyMap[nums[i] + 1] ; ok {
			uf.union(i,keyMap[nums[i] + 1])
		}
		keyMap[nums[i]] = i
	}
	return uf.getMaxConnectionSize()
}

type UF struct {
	parent []int
	size []int
}
//查找这个数组下标n的爹
func  newUf(n int) *UF {
	parent := make([]int,n)
	size := make([]int,n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return  &UF{
		 parent,
		 size,
	}
}
func (uf *UF) getMaxConnectionSize() int {
	max := math.MinInt64
	for i := 0; i < len(uf.size); i++ {
		if uf.size[i] > max {
			max = uf.size[i]
		}
	}
	return max
}

func (uf *UF) union(x,y int)  {
	if uf.parent[x] == uf.parent[y] {
		return
	}
	parentX := uf.find(x)
	parentY := uf.find(y)
	if uf.size[parentX] > uf.size[parentY] {
		uf.parent[parentY] = parentX
		uf.size[parentX] += uf.size[parentY]
	}else {
		uf.parent[parentX] = parentY
		uf.size[parentY] += uf.size[parentX]
	}
}
func (uf *UF) find(x int) int {
	if x == uf.parent[x] {
		return x
	}else {
	   return uf.find(uf.parent[x])
	}
}


func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	longestStreak, curStreak := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				curStreak++
			} else {
				longestStreak = max(longestStreak, curStreak)
				curStreak = 1
			}
		}
	}

	return max(longestStreak, curStreak)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
