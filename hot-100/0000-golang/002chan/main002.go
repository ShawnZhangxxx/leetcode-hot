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
	go func() {
		i:= 0
		for {
			i++
			fmt.Println("new goroutine" ,i)
			time.Sleep(time.Second)
		}
	}()

	i:=0
	for  {
		i++
		fmt.Println("main goroutine",i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}

	}
}
