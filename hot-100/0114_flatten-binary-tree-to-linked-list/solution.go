/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-06-28 17:04:57
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-06-28 17:08:30
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
	//       3
	//   9	    	20
	// `         15     7
	preorder := initTreeLayer(3,9,20,15,7)
	res := flattern2(preorder)
	printTreeLayer(res)
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
//归并排序-后序遍历思路
func flattern2(root *TreeNode) *TreeNode{
	if root == nil {
		return nil
	}
	//fmt.Println(root.Val)
	left := flattern2(root.Left)
	right := flattern2(root.Right)
	head:=root
	root.Right = left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
	return head
}


func flattern(root *TreeNode) {
	if root == nil {
		return
	}

	flattern(root.Left)
	flattern(root.Right)
	left, right := root.Left, root.Right
	root.Left = nil
	root.Right = left
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	cur.Right = right
}
