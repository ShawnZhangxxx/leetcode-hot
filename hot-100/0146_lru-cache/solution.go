/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-04 15:09:01
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-04 15:40:51
 * @Description:
 */

package main

import (
	"fmt"
)

func main() {
	lruCache := Constructor2(5)
	fmt.Println(lruCache)
}

type Node2 struct {
	key, value int
	next       *Node2
	prev       *Node2
}

type LRUCache2 struct {
	size, capacity int
	cache          map[int]*Node2
	head, tail     *Node2
}

func Constructor2(capacity int) LRUCache2 {
	lruCache := LRUCache2{
		size:     0,
		capacity: capacity,
		head:     &Node2{},
		tail:     &Node2{},
		cache:    make(map[int]*Node2, 0),
	}
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache

}

func (this *LRUCache2) deleteNode2(node2 *Node2) {
	node2.prev.next = node2.next
	node2.next.prev = node2.prev
}
func (this *LRUCache2) addToTail2(node2 *Node2) {
	this.tail.prev.next = node2
	node2.next = this.tail
}
func (this *LRUCache2) updateNode2(node2 *Node2) {
	this.addToTail2(node2)
	this.deleteNode2(node2)
}

func (this *LRUCache2) Get2(number int) int {
	if v, ok := this.cache[number]; ok {
		this.updateNode2(v)
		return v.value
	}
	return -1
}

func (this *LRUCache2) Put2(number int) int {

}

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*Node
	head, tail     *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*Node, 0),
		head:     &Node{},
		tail:     &Node{},
	}

	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) addToTail(node *Node) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) moveToTail(node *Node) {
	this.removeNode(node)
	this.addToTail(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		node.value = value
		return
	}

	if this.size == this.capacity {
		this.size--
		node := this.head.next
		this.removeNode(node)
		delete(this.cache, node.key)
	}

	node := &Node{key: key, value: value}
	this.size++
	this.addToTail(node)
	this.cache[key] = node
}
