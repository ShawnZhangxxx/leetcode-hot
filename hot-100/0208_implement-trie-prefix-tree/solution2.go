/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-10 11:01:26
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-10 11:08:14
 * @Description:
 */

package main

import "fmt"

type  Tire  struct{
	Next map[rune]*Tire
	IsWord bool
}
func Constructor() Tire {
	return Tire{
		Next : make(map[rune]*Tire),
		IsWord :false,
	}
}

func (this *Tire) Search (word string ) bool {
	for _, v := range word {
		if this.Next[v] == nil {
			return false
		}else {
			this = this.Next[v]
		}
	}
	if this.IsWord == false {
		return false
	}else {
		return true
	}
}
func (this *Tire) StartsWith (word string ) bool {
	for _, v := range word {
		if this.Next[v] == nil {
			return false
		}else {
			this = this.Next[v]
		}
	}

		return true

}
func (this *Tire) Insert (word string)  {
	for _, v := range word {
		if this.Next[v] == nil {
			this.Next[v] = &Tire{
				IsWord: false,
				Next: make(map[rune]*Tire),
			}
		}
		this = this.Next[v]
	}
	this.IsWord = true
}
func  main()  {
	tire :=	Constructor()
	fmt.Println(tire)
	tire.Insert("abc")
	a := tire.Search("bcd")
	b := tire.StartsWith("abd")
	fmt.Println(tire)
	fmt.Println(a)
	fmt.Println(b)

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
