package main

import (
	"fmt"

	"github.com/Onlymiind/DataStructures/list"
)

func add(lhs int, rhs int) int {
	return lhs + rhs
}

func addPrint(lhs int, rhs int) int {
	fmt.Println(lhs, " + ", rhs)
	return lhs + rhs
}

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1;
	} else { 
		return Fibonacci(n-1) + Fibonacci(n-2);
	}
}

func main() {
	l := list.NewIntList()
	for i := 0; i <= 100; i++ {
		l.PushBack(i)
	}
	fmt.Println(l.String())
}