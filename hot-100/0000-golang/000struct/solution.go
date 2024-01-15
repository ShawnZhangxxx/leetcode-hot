/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 14:02'
*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {
  structA :=	struct {
	  a int8
	  b int64
	  c int32
	  d int16
	}{}
	structB :=	struct {
		a int8
		b int16
		c int32
		d int64
	}{}
  fmt.Println(unsafe.Sizeof(structA))
  fmt.Println(unsafe.Sizeof(structB))
}

