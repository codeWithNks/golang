package main

import "fmt"

func main() {
	fmt.Printf("Learning Arrays and Slices in Go\n")
	// Arrays====================================================================
	var arr [5]int
	fmt.Println(arr)
	var arr1 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	var arr2 = [5]int{1, 2, 3}
	fmt.Println(arr2)
	// length of all 3
	fmt.Println(len(arr))
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))

	// Slices====================================================================
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(slice[1:3])
	fmt.Println(slice[:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[:])
	// append
	slice = append(slice, 6)
	fmt.Println(slice)
	// copy
	slice1 := make([]int, 5) // Another way to create slice
	slice1[0] = 11
	slice1[1] = 22
	slice1[2] = 33
	slice1[3] = 44
	slice1[4] = 55

	//slice1[5] = 6 // This will give error
	slice1 = append(slice1, 6) // This will work
	fmt.Println(slice1)

	copy(slice1, slice) // overwrites slice1 with slice

	fmt.Println(slice1)
	slice = append(slice[:2], slice[3:]...) // delete from slice
	fmt.Println(slice)
	// Difference between array and slice parallely
	// 1. Array is fixed size, slice is dynamic
	// 2. Array is value type, slice is reference type
	// 3. Array is passed by value, slice is passed by reference

}
