package main

import (
	"errors" // library for errors variables
	"fmt"
)

func main() {
	printMe("Hello")

	numerator, denominator := 11, 2

	var result, remainder, err = intDivision(numerator, denominator)

	fmt.Printf("The result of the integer division is %v wth remainder %v\n", result, remainder) // printf funciton we can format the output
	// %v is the general identifier instead of %s, %d in c

	// what is 0 was given as the denominator
	numerator, denominator = 11, 0
	result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error() + "\n")
	} else if remainder == 0 { // && and || can be used with similar meaning as in C and C++
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v wth remainder %v\n", result, remainder)
	}

	switch { // break after each case is already implied and need not be written
	case err != nil:
		fmt.Printf(err.Error() + "\n")
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer division is %v wth remainder %v\n", result, remainder)
	}

	// runtime error because division by 0
	// go gievs a run time error, but it is handeled using a error variable called err
}

func printMe(printValue string) {
	// fmt.Println("Hello world!")
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) { // this funciton can return multiple values!!!!!!!!!!!!!!!!!
	var err error
	// by deafault the value of err is nil
	if denominator == 0 {
		err = errors.New("Cannot devide by Zero") // configures an error message
		return 0, 0, err                          // returns an error message
	}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
