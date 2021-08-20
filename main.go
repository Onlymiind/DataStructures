package main

import (
	"fmt"

	"github.com/Onlymiind/DataStructures/list"
)

func main() {
	var x int = 10
	var i interface{} = x
	var j interface{} = x

	l := list.NewIntList()
	l.PushBack("Hey")

	x = 11
	fmt.Println(x, i, j)
}
