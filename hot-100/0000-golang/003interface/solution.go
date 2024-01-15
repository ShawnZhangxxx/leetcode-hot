/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import "fmt"

type People interface {
	Speak(string) string
}
type Student struct{}

 func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
		 } else {
		 talk = "您好"
		 }
 		return
 }

 func main() {
	 var peo People = &Student{}
	 think := "bitch"
	 fmt.Println(peo.Speak(think))
	 }
