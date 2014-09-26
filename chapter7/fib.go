package main

import "fmt"

func main() {

	fmt.Println(fib(10))
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
