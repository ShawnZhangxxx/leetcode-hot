/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-18 14:02:05
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-18 14:04:30
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
	//       10
	//   7	    	8
	// 3   11               9
	var root *TreeNode
	root = initTreeLayer(10,7,8,3)
	printTreeLayer(root)
	node1 := &TreeNode{9,nil,nil}
	node2 := &TreeNode{11,nil,nil}
	root.Right.Right = node1
	root.Right.Left = node2
	res := lowestCommonAncestor2(root,node1,node2)
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
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || p == root || q == root {
		return root
	}
	l := lowestCommonAncestor2(root.Left,p,q)
	r := lowestCommonAncestor2(root.Right,p,q)

	if l != nil && r != nil {
		return root
	}else if l == nil {
		return r
	}else {
		return l
	}

}
