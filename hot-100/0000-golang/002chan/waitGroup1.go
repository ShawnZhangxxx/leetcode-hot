package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int)  {
	defer wg.Done()
	fmt.Println("hello groutine!" , i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait() //等待登记go routine 都结束
}