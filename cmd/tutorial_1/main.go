package main

import (
	"fmt"          // used for input and output, like stdio or iostreaem
	"unicode/utf8" // contains tools to get the number of characters in a string
)

func main() {

	// ; is optional at every line
	// every new line acts like ; so ; can be skipped if the next command is on a new line
	// :-)

	var intNum uint16 = 32767 + 1
	fmt.Println(intNum)

	// other options are:
	// int8, int16, int32, int64 and also uint8, uint 16, uint32, uint64
	// the default value is 0

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)
	// only 2 options: float32 or float64 has to be specified, float gives an error

	// var result float32 = floatNum + intNum
	// cannot add different types of data.
	// should convert data types into a common format to perform arthemetics
	var result float32 = floatNum + float32(intNum)
	fmt.Println(result)

	var intNum1 = 3
	var intNum2 = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = `Hello 
	World`
	// here the above string is equivalent to "Hello \nWorld"
	// thus one can use backquoties `` we can format the string directely
	myString = "Hello" + " " + "World"
	fmt.Println(myString)

	// length of a string
	fmt.Println(len(myString))
	// len function gives out the length in bytes
	// go uses UTF 8 encoding, so characters out of the vanilla ASCII character set are storeed with more than a single byte

	fmt.Println(len("γ")) // output is 2

	// utf8.RuneCountInString()gives us the number of characters in the string
	fmt.Println(utf8.RuneCountInString("γ")) // output is 1

	// rune are a seperate data type in go and represent characters

	var myRune rune = 'a' // single quotes
	fmt.Println(myRune)   // prints utf 8 integer value i.e 97 in this case

	var myBoolean bool = false
	fmt.Println(myBoolean) // prints false

	// default values
	// integers = 0
	// strings = ""
	// rune = 0
	// bool = false

	var myVar = "text" // here the type is infered and myVar is now a string variable
	fmt.Println(myVar)

	// we can drop the var keyword and use the shorthand := instead

	myVar2 := "text"
	fmt.Println(myVar2)

	// initialize multiple variables

	var var1, var2 = 1, 2
	// similar to var var1, var2 int = 1, 2
	// similar to var1, var2 := 1, 2
	fmt.Println(var1, var2)

	// ADDING THE TYPE WHEN IT ISNTT OBVIOUS IS GOOD PRACTICE, can help deal with functions whose return value is unknown

	// CONSTANTS

	// constants are same as vairables but you cannot changne its value once defines

	const myConst string = "const value"
	// myConst = "another value" this given an error
	// also you cannot declate constats, we have to initialist it's values
	fmt.Println(myConst)

	// const is helpful when using universal constants like PI = 3.14

	const PI = 3.14

}
