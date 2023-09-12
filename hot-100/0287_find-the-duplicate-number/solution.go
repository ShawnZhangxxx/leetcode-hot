/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:16:39
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 10:30:29
 * @Description:
 */

package main

func main()  {

}



func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return slow
}
