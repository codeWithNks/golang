package main

import "fmt"

//constants
const Pi = 3.14 // capital letter means it is exported or public
const pi = 3.14 // small letter means it is private

func main() {
	var name string = "Nks"
	fmt.Printf("Variable is of type %T\n", name)

	var isTrue bool = true
	fmt.Printf("Variable is of type %T\n", isTrue)

	var smallValue uint8 = 255
	fmt.Printf("Variable is of type %T\n", smallValue)

	var smallFloat float32 = 255.455
	fmt.Printf("Variable is of type %T\n", smallFloat)

	var largeFloat float64 = 255.455454565
	fmt.Printf("Variable is of type %T\n", largeFloat)

	//default values and aliases
	var defaultInt int
	fmt.Println("Default value of int is ", defaultInt)

	var defaultString string
	fmt.Println("Default value of string is ", defaultString)

	//implicit type
	var implicitInt = 10
	fmt.Printf("Variable is of type %T\n", implicitInt)

	//no var style using operator := only works inside functions
	implicitString := "Nks"
	fmt.Println("value of implicitString is ", implicitString)

	//constants
	fmt.Println("Value of Pi is ", Pi)
	fmt.Println("Value of pi is ", pi)

}
