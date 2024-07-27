package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("hey, enter your age: ")
	input := bufio.NewReader(os.Stdin)
	age, _ := input.ReadString('\n')
	fmt.Println("Your age is ", age)
	fmt.Printf("Variable is of type %T\n", age)

	// conversion
	newAge, err := strconv.ParseFloat(strings.TrimSpace(age), 64) // TrimSpace removes leading and trailing spaces
	if err != nil {
		fmt.Println("Error in conversion", err)
	} else {
		fmt.Println("Current age is ", newAge)
		fmt.Println("Age after 5 years is ", newAge+5)
	}

}
