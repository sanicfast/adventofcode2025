package main

import (
	"fmt"
)

func main() {
	// Declare a variable and assign the binary literal for 10
	var myBinaryValue int = 0b1010

	// Declare a constant and assign the binary literal for 10
	const anotherBinaryValue int = 0b1010
	var2 := myBinaryValue ^ anotherBinaryValue
	fmt.Println(var2, "var2")
	fmt.Printf("var2 (binary): %b\n", var2)

	// Print the decimal value
	fmt.Println("myBinaryValue (decimal):", myBinaryValue)
	fmt.Println("anotherBinaryValue (decimal):", anotherBinaryValue)

	// Print the binary representation using the %b format specifier
	fmt.Printf("myBinaryValue (binary): %b\n", myBinaryValue)
	fmt.Printf("anotherBinaryValue (binary): %b\n", anotherBinaryValue)
}
