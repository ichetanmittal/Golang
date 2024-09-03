package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointers")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var myNumberPtr = &myNumber
	fmt.Println("Value of pointer is ", myNumberPtr)
	fmt.Println("Value of pointer is ", *myNumberPtr)

	*myNumberPtr = *myNumberPtr * 2
	fmt.Println("New value is: ", myNumber)
}
