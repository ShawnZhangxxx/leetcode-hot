/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int8,5)

	go go2(chan1)
	go go1(chan1)
	time.Sleep(10)
}
func go2(chan1 chan int8)  {
	for i := 0; i < 5; i++ {
		chan1 <- 1
	}
	close(chan1)
}
func go1(chan1 chan int8)  {
	for  {
		select {
		case <-chan1 :
			fmt.Println(1)
		case chan1 <-1:
			fmt.Println(2)
		default:
			
		}
	}
}
