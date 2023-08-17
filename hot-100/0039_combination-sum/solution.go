/*
__author__ = 'robin-luo'
__date__ = '2023/03/13 19:15'
*/

package main

import "fmt"

//重点题 TODO
func main()  {

	ans := combinationSum2([]int{1,2,4,5,10},10)
	fmt.Println(ans)
}

func combinationSum2(candidates []int,target int)(ans [][]int)  {
	combine := []int {}
	var dfs func(target,idx int)   //如果是def := func(){} 的形式不能递归调用
	dfs = func (target int,index int) {
			if index == len(candidates) {
				return
			}
			if target == 0 {

				ans  = append(ans,append([]int(nil),combine...)) //这么写的目的 是防止被引用,你append的只是地址,
				//ans  = append(ans,combine)
				return
			}
			dfs(target,index+1)
			if target-candidates[index] >= 0  {//选择当前元素
				combine = append(combine,candidates[index])
				dfs(target-candidates[index],index)
				combine = combine[:len(combine)-1] //这一步是回溯的核心
			}
		}
		dfs(target,0)
	return ans
}


//var candidates []int = []int{1,5,10}
//var ans  [][]int = [][]int{}
//
//func dfs (target int,combine []int,index int) {
//	if index == len(candidates) {
//		return
//	}
//	if target == 0 {
//		ans  = append(ans,combine)
//		return
//	}
//	dfs(target,combine,index+1)
//	if target > 0  {//选择当前元素
//		combine = append(combine,candidates[index])
//		dfs(target-candidates[index],combine,index)
//	}
//}





func combinationSum(candidates []int, target int) [][]int {
	var (
		track []int
		res   [][]int
	)
	backTracking(target, 0, 0, candidates, track, &res)
	return res
}




func backTracking(target, sum, startIndex int, candidates, track []int, res *[][]int) {
	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	if sum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backTracking(target, sum, i, candidates, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}
