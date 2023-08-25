/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-27 10:13:43
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 10:37:25
 * @Description:
 */
/*
__author__ = 'robin-luo'
__date__ = '2023/06/27 10:13'
*/

package main

import "fmt"

func main()  {
	//res:=largestRectangleArea2([]int{6,7,5,2,4,5,9,3})
	res:=largestRectangleArea2([]int{2,1,5,6,2,3})
	fmt.Println(res)
}

//单调栈,
func largestRectangleArea2(heights []int)int  {
	var left = make([]int,len(heights))
	var right  = make([]int,len(heights))
	var max int
	var stack []int
	for i := 0; i < len(heights); i++ {
		for len(stack) >= 0 {
			if len(stack) == 0 {//== 0 压栈,left 加-1
				left[i] = -1
				stack = append(stack,i)
				break
			}else if 	heights[stack[len(stack)-1]] < heights[i] {//左侧比当前小,压栈,压left
				left[i] = stack[len(stack)-1]
				stack = append(stack,i)
				break
			}else {//左侧大,弹栈,继续比较,直至stack = 0
				stack = stack[:len(stack)-1]
			}
		}
	}
	stack = []int{} //清空栈
	fmt.Println(left)
	for i := len(heights) -1; i > -1; i-- {
		for len(stack) >= 0 {
			if len(stack) == 0 {//== 0 压栈,left 加-1
				right[i] = len(heights)
				stack = append(stack,i)
				break
			}else if heights[stack[len(stack)-1]] < heights[i] {//右侧比当前小,压栈,压left
				right[i] = stack[len(stack)-1]
				stack = append(stack,i)
				break
			}else {//左侧大,弹栈,继续比较,直至stack = 0
				stack = stack[:len(stack)-1]
			}
		}
	}

	var area int
	for i := 0; i < len(heights) ; i++ {
		area = (right[i] - left[i] - 1) * heights[i]
		if max < area {
			max = area
		}
	}

	return max

}






func largestRectangleArea(heights []int) int {
	var (
		stack  []int
		i, res = 0, 0
	)

	for i < len(heights) {
		if len(stack) == 0 || heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}
			res = max(res, width*heights[top])
			i--
		}
		i++
	}

	for len(stack) != 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		var width int
		if len(stack) == 0 {
			width = i
		} else {
			width = len(heights) - stack[len(stack)-1] - 1
		}

		res = max(res, width*heights[top])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
