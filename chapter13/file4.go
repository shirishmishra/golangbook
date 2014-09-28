package main

import (
		"fmt"
		"os"
)

func main() {
	dir, err := os.Open("..")

	if err != nil {
		return
	}
	
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)

	for _, f := range fileInfos {
		fmt.Println(f.Name())
	}
}
