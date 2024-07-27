package main

import "fmt"

func main() {
	fmt.Printf("Learning Pointers in Go\n")
	var ptr *int
	fmt.Printf("Value of ptr: %v\n", ptr) // nil
	fmt.Printf("Address of ptr: %v\n", &ptr)

	// %v is used to print the value in default format
	myNumber := 10
	ptr = &myNumber                                    // Assigning address of myNumber to ptr
	fmt.Printf("Value of myNumber: %v\n", myNumber)    // Value of myNumber
	fmt.Printf("Value at address of ptr: %v\n", *ptr)  // Dereferencing ptr (* means value at address)
	fmt.Printf("Address of myNumber: %v\n", &myNumber) // Address of myNumber (& means address of)
	fmt.Printf("Value of ptr: %v\n", ptr)              // Address of myNumber (ptr holds address of myNumber)
	fmt.Printf("Address of ptr: %v\n", &ptr)           // Address of ptr

	// Changing value of myNumber using ptr
	fmt.Printf("=====================================\n")
	*ptr = 20                                          // *ptr is same as myNumber
	fmt.Printf("Value of myNumber: %v\n", myNumber)    // Value of myNumber
	fmt.Printf("Value at address of ptr: %v\n", *ptr)  // Dereferencing ptr
	fmt.Printf("Address of myNumber: %v\n", &myNumber) // Address of myNumber
	fmt.Printf("Value of ptr: %v\n", ptr)              // Address of myNumber
	fmt.Printf("Address of ptr: %v\n", &ptr)           // Address of ptr

}
