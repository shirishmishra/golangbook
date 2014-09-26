package main

import "fmt"

func main() {
	fmt.Println(halve(1))
	fmt.Println(halve(2))
}

func halve(val int) (int, bool) {
	return val / 2, val % 2 == 0
}
