package main

import (
	"custom_list"
	"fmt"
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
	var a, b int
	var l custom_list.List
	l.PushBack(1)
	plus := add
	fmt.Scanf("%d", &a);
	fmt.Scanf("%d", &b);
	fmt.Println(plus(a, b));
	plus = addPrint
	fmt.Print(plus(a,b))
}
