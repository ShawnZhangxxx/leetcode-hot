/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-24 10:36:06
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-24 10:43:43
 * @Description:
 */

package main

import (
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
	res []string
}

func Constructor() Codec {
	return Codec{}
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
	this.res = strings.Split(data, ",")
	return this.dfs()
}

func (this *Codec) dfs() *TreeNode {
	node := this.res[0]
	this.res = this.res[1:]
	if node == "#" {
		return nil
	}

	val, _ := strconv.Atoi(node)
	root := &TreeNode{
		Val:   val,
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
