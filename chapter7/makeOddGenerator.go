package main

import "fmt"

func makeOddGenerator() func() int {
	i := -1

	return func() int {
		i += 2

		return i
	}

}
func main() {
		nextOddGenerator := makeOddGenerator()

		fmt.Println(nextOddGenerator())
	    fmt.Println(nextOddGenerator())
		fmt.Println(nextOddGenerator())
		
}
