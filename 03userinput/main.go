package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")
	// comma ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ", input)
}
