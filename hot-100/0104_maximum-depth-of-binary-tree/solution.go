/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 15:45:19
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 15:58:11
 * @Description:
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func main() {
	//       5
	//   3	    	7
	// 2   4
	var root *TreeNode
	root = initTreeLayer(10,8,8,7,7,7,7)
	printTreeLayer(root)
	res := maxDepth2(root)
	fmt.Println(res)
}
//初始化一个搜索二叉树
func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}
	return root

}
func initTreeLayer( Data ...int) *TreeNode {
	var queue []*TreeNode
	head := &TreeNode{
		Val: Data[0],
	}
	queue =  append(queue,head)
	i := 0
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		i = i+1
		if i<len(Data) {
			top.Left = &TreeNode{Val: Data[i]}
			queue = append(queue,top.Left)
		}else {
			break
		}
		i = i+1
		if i<len(Data) {
			top.Right = &TreeNode{Val: Data[i]}
			queue = append(queue,top.Right)
		}else {
			break
		}
	}
	return  head
}
//层序遍历
func printTreeLayer( root *TreeNode)  {
	var queue []*TreeNode
	queue =  append(queue,root)
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Println(top.Val)
		if top.Left != nil {
			queue = append(queue,top.Left)
		}
		if top.Right != nil {
			queue = append(queue,top.Right)
		}
	}
}
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth2(root.Left)+1,maxDepth2(root.Right)+1)
}
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// 迭代
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	q := []*TreeNode{root}
	for len(q) != 0 {
		size := len(q)
		level++
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return level
}
