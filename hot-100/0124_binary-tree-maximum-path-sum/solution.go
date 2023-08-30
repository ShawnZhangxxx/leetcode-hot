/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-03 10:10:29
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-03 11:04:16
 * @Description:
 */

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//       3
	//   9	    	20
	// 15   7
	preorder := initTreeLayer(3,9,20,15,7)
	max := maxPathSum2(preorder)
	fmt.Println(max)
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

func maxPathSum2(root *TreeNode) int {
	maxSum := math.MinInt64
	dfs2(root,&maxSum)
	return maxSum
}

func dfs2(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, maxSum)
	r := dfs(root.Right, maxSum)
	*maxSum = max(*maxSum, root.Val+l+r, root.Val+l, root.Val+r, root.Val)
	return max(root.Val, root.Val+l, root.Val+r)
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	dfs(root, &maxSum)
	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, maxSum)
	r := dfs(root.Right, maxSum)
	*maxSum = max(*maxSum, root.Val+l+r, root.Val+l, root.Val+r, root.Val)
	return max(root.Val, root.Val+l, root.Val+r)
}

func max(nums ...int) int {
	res := math.MinInt64
	for _, item := range nums {
		if item > res {
			res = item
		}
	}

	return res
}
