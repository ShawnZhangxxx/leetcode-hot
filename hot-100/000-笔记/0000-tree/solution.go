/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
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
	root = initTreeLayer(10,8,8,7,7,7,7)
	printTreeLayer(root)
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