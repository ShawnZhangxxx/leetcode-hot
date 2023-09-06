///*
// * @Author: lawtech0902 584563542@qq.com
// * @Date: 2023-07-10 11:01:26
// * @LastEditors: lawtech0902 584563542@qq.com
// * @LastEditTime: 2023-07-10 11:08:14
// * @Description:
// */
//
package main
//
//type Trie struct {
//	next   map[rune]*Trie
//	isWord bool
//}
//
//func Constructor() Trie {
//	return Trie{
//		next:   make(map[rune]*Trie),
//		isWord: false,
//	}
//}
//
//func (this *Trie) Insert(word string) {
//	for _, v := range word {
//		if this.next[v] == nil {
//			this.next[v] = &Trie{
//				next:   make(map[rune]*Trie),
//				isWord: false,
//			}
//		}
//
//		this = this.next[v]
//	}
//
//	this.isWord = true
//}
//
//func (this *Trie) Search(word string) bool {
//	node := this.prefixNode(word)
//	return node != nil && node.isWord
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	node := this.prefixNode(prefix)
//	return node != nil
//}
//
//func (this *Trie) prefixNode(prefix string) *Trie {
//	for _, v := range prefix {
//		if this.next[v] == nil {
//			return nil
//		}
//
//		this = this.next[v]
//	}
//
//	return this
//}
//
///**
// * Your Trie object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Insert(word);
// * param_2 := obj.Search(word);
// * param_3 := obj.StartsWith(prefix);
// */
