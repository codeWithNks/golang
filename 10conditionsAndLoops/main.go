package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Learning about conditions and loops in Go")

	// 1.if-else==================================================
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is equal to 5")
	}

	// 1.1 if-else with initialization
	if y := 10; y > 5 {
		fmt.Println("y is greater than 5")
	} else if y < 5 {
		fmt.Println("y is less than 5")
	} else {
		fmt.Println("y is equal to 5")
	}

	// 2. switch-case==============================================
	// 2.1 switch-case with default
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	case 3:
		fmt.Println("x is 3")
		fallthrough // fallthrough is used to execute the next case as well
	default:
		fmt.Println("x is not 1, 2 or 3")
	}

	// 2.2 switch-case with multiple cases
	switch x {
	case 1, 2, 3:
		fmt.Println("x is 1, 2 or 3")
	default:
		fmt.Println("x is not 1, 2 or 3 in multiple cases")
	}

	// 2.3 switch-case with expression
	switch { // here we are not passing any value to switch, so it will check the expression in case
	case x > 5:
		fmt.Println("x is greater than 5 in expression")
	case x < 5:
		fmt.Println("x is less than 5 in expression")
	default:
		fmt.Println("x is equal to 5 in expression")
	}

	// 3. for loop=================================================
	// 3.1 for loop with initialization, condition and post
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// 3.2 for loop with only condition
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 3.3 for loop with range
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Println("Index:", index, "Value:", value)
	}

	// 3.4 for loop with range and only value
	for _, value := range arr { // here we are not using index, so we are using _ to ignore it and _ is a blank identifier
		fmt.Println("Value:", value)
	}

	// 4. break and continue=======================================
	// 4.1 break
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// 4.2 continue
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// 5. goto=====================================================
	// 5.1 goto
	i := 0
Loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto Loop
	}

	// Additional--> Random number generation between 1-6 using rand
	// 6. Random number generation between 1-6 using rand

	rand.Seed(time.Now().UnixNano()) // seeding with times means it will generate random number every time
	fmt.Println(rand.Intn(6) + 1)    // rand.Intn(6) will generate random number between 0-5, so adding 1 to make it between 1-6

}
