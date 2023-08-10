/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-26 15:41:55
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-26 15:45:31
 * @Description:
 */

package main

func hanmingDistance(x, y int) int {
	xor := x ^ y
	count := 0
	for xor != 0 {
		if xor&1 == 1 {
			count++
		}

		xor >>= 1
	}

	return count
}
