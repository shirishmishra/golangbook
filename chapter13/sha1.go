package main

import (
		"fmt"
		"crypto/sha1"
)

func main() {
	h := sha1.New()

	h.Write([]byte("test"))

	v := h.Sum([]byte{})

	fmt.Println(v)

}
