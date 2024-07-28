package main

import "fmt"

func main() {
	fmt.Println("Learning about functions in Go")

	// 1. Function without parameters and return type
	greet()

	// 2. Function with parameters and return type
	sum := add(10, 20)
	fmt.Println("Sum of 10 and 20 is", sum)

	// 3. Function with multiple return values
	quotient, remainder := divide(100, 7)
	fmt.Println("Quotient is", quotient, "and remainder is", remainder)

	// 4. Function with named return values
	quotient, remainder = divideWithNamedReturnValues(100, 7)
	fmt.Println("Named Quotient is", quotient, "and Named remainder is", remainder)

	// 5. Function with variadic parameters
	sum = sumAll(10, 20, 30, 40, 50)
	fmt.Println("Sum of 10, 20, 30, 40 and 50 is", sum)

}

// greet is a function without parameters and return type
func greet() {
	fmt.Println("Hello, World!")
}

// add is a function with parameters and return type
func add(a int, b int) int {
	return a + b
}

// divide is a function with multiple return values
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// divideWithNamedReturnValues is a function with named return values
func divideWithNamedReturnValues(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // return statement without any value will return the named return values
}

// sumAll is a function with variadic parameters
func sumAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
