package main

import "fmt"

func main() {
	// Structs====================================================================
	var u User
	u.FirstName = "John"
	u.LastName = "Doe"
	u.Age = 30
	u.Email = "nks@yopmail.com"
	fmt.Println(u)

	// Printing with field names using +v
	fmt.Printf("%+v\n", u) // %+v prints field names as well as values in struct

	// Another way to create struct
	u1 := User{FirstName: "Jane", LastName: "Doe", Age: 25, Email: "nks@yopmail.com"}
	fmt.Printf("%+v\n", u1)

	// Another way to create struct
	u2 := User{"Jack", "Doe", 35, "nks3@yopmail.com"}
	fmt.Printf("%+v\n", u2)

	// Anonymous struct  -  struct without a name - used for one time use case
	u3 := struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}{
		FirstName: "Jill",
		LastName:  "Doe",
		Age:       40,
		Email:     "xyz@gmail.com",
	}

	fmt.Printf("%+v\n", u3)
}

type User struct { //Capitalized User to make it public
	FirstName string //Capitalized FirstName to make it public
	LastName  string
	Age       int
	Email     string
}
