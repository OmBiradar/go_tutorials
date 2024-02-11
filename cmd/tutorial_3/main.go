package main

import (
	"fmt"
)

func main() {
	printMe("Hello")

	numerator, denominator := 11, 2

	var result, remainder = intDivision(numerator, denominator)
	fmt.Printf("The result of the integer division is %v wth remainder %v\n", result, remainder) // printf funciton we can format the output
	// %v is the general identifier instead of %s, %d in c

	// what is 0 was given as the denominator
	fmt.Println(intDivision(11, 0)) // runtime error because division by 0
	// go gievs a run time error
}

func printMe(printValue string) {
	// fmt.Println("Hello world!")
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) { // this funciton can return multiple values!!!!!!!!!!!!!!!!!
	var err error
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder
}