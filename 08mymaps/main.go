package main

import "fmt"

func main() {
	fmt.Println("Learning Maps in Go")
	// Maps====================================================================
	var m map[string]int
	fmt.Println(m)
	//Using make function
	m = make(map[string]int)
	fmt.Println(m)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3
	fmt.Println(m)
	fmt.Println(m["a"])
	// delete key from map
	delete(m, "a")
	fmt.Println(m)
	// check if key exists in map
	val, exists := m["a"] //  *** Important :  val will be 0 if key doesn't exist
	fmt.Println(val, exists)

	// Iterating over map
	for key, value := range m {
		fmt.Println(key, value)
	}
	// Difference between map and slice parallely
	// 1. Slice is ordered, map is unordered
	// 2. Slice is indexed, map is key-value pair
	// 3. Slice is dynamic, map is dynamic
	// 4. Slice is passed by reference, map is passed by reference
}
