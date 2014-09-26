package main

import "fmt"

func main() {
	fmt.Print("Enter temperature in celsius: ")
	
	var temp float64

	fmt.Scanf("%f", &temp)

	fmt.Println("In Fahrenheit:", (1.8 * temp + 32)) 
}
