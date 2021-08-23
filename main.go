package main

import (
	"fmt"

	"github.com/Onlymiind/DataStructures/list"
)

func main() {
	l := list.NewList()
	for i := 0; i < 2; i++ {
		l.PushBack("Hey")
	}
	fmt.Println(l)
}
