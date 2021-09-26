package main

import (
	"fmt"
	"testing"
)

func Test_Stack(t *testing.T) {
	stack := Init(10)
	for i := 0; i < 10; i++ {
		stack.Push(i)
		//fmt.Println(stack.Pop())
	}

	fmt.Println(stack.Pop())

}
