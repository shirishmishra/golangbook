package main

import "fmt"

func main() {
	fmt.Println(max(1,2,3,4))
	fmt.Println(max(17,1,23,221,1322,334))
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return -1
	}
	
	max := vals[0]

	for _,v := range vals {
		if v > max {
			max = v
		}
	}

	return max
}
