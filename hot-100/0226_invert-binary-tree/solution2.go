/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-18 10:38:35
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-18 13:41:25
 * @Description:
 */

package main

import "fmt"

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
	root = initTreeLayer(10,9,8,6,2,7,7)
	printTreeLayer(root)
	invertTree2(root)
	printTreeLayer(root)
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
	for len(queue) > 0 { //
		size := len(queue)
		for i := 0; i < size; i++ { //这是一层
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
}

//翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left,root.Right = invertTree(root.Right),invertTree(root.Left)
	return  root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		 return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		top.Left,top.Right = top.Right,top.Left
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return root


}