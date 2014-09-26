package main

import "fmt"

func main() {
	fmt.Print("Enter measurement in feet: ")

	var feet float64

	fmt.Scanf("%f", &feet)
	
	fmt.Println("Measurement in meters:", (feet * 0.3048))
}
