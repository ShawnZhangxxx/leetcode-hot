/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:36:06
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 10:43:43
 * @Description:
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Codec struct {
	str []string
}

func Constructor() Codec {
	return Codec{}
}
func main() {
	//       5
	//   3	    	7
	// 2   4
	var root *TreeNode
	root = initTreeLayer(10,8,8,7,7,7,7)
	ser := Constructor();
	deser := Constructor();

	data := ser.serialize(root);
	ans := deser.deserialize(data);
	printTreeLayer(ans)
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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.str = strings.Split(data,",") //数组的 先序遍历
	return this.dfs()

}
//将树变成字符串
func (this *Codec) dfs() *TreeNode {
	node := this.str[0]
	this.str = this.str[1:]
	if node == "#" {
		return nil
	}
	val,_ := strconv.Atoi(node)
	root := &TreeNode{
		Val: val,
		Left:  this.dfs(),
		Right: this.dfs(),
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
