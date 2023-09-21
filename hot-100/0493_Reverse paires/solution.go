/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 15:51:33
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 09:37:27
 * @Description:
 */

package main

import "fmt"

//Given an integer array nums, return the number of reverse pairs in the array.
//
//
// A reverse pair is a pair (i, j) where:
//
//
// 0 <= i < j < nums.length and
// nums[i] > 2 * nums[j].
//
//
//
// Example 1:
//
//
//Input: nums = [1,3,2,3,1]
//Output: 2
//Explanation: The reverse pairs are:
//(1, 4) --> nums[1] = 3, nums[4] = 1, 3 > 2 * 1
//(3, 4) --> nums[3] = 3, nums[4] = 1, 3 > 2 * 1
//
//
// Example 2:
//
//
//Input: nums = [2,4,3,5,1]
//Output: 3
//Explanation: The reverse pairs are:
//(1, 4) --> nums[1] = 4, nums[4] = 1, 4 > 2 * 1
//(2, 4) --> nums[2] = 3, nums[4] = 1, 3 > 2 * 1
//(3, 4) --> nums[3] = 5, nums[4] = 1, 5 > 2 * 1
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 5 * 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
// Related Topics 树状数组 线段树 数组 二分查找 分治 有序集合 归并排序 👍 419 👎 0

func main(){
	res := reversePairs([]int{1,3,2,3,1 })
	fmt.Println(res)
}
func reversePairs(nums []int) int {

}
