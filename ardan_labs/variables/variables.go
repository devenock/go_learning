package main

import "fmt"

func main() {
	// variables are always the heart of a programming language as they provide the ability to read from and write to memory

	// zero value: every variable you declare without assigning it a value has a default value known as zero value.
	// the zero value concept exist in Go to enforce data integrity and it is not free.

	// Type					Zero Value
	// Boolean				false
	// Integer				0
	// Float 				0
	// complex 				0i
	// String				""
	// Pointer				nil

	// declaring variables
	// You can either use the (var) kyword or use the Walrus symbol(:=) to declare variable
	var name string
	var phoneNo int
	var isComing bool
	var secondName *string
	var average float64

	// check the types
	fmt.Printf("here are the variable types: %T and value: %v\n", name, name)
	fmt.Printf("here are the variable types: %T and value: %v\n", phoneNo, phoneNo)
	fmt.Printf("here are the variable types: %T and value: %v\n", isComing, isComing)
	fmt.Printf("here are the variable types: %T and value: %v\n", secondName, secondName)
	fmt.Printf("here are the variable types: %T and value: %v\n", average, average)

	// separator
	fmt.Println("<==========================SHORT HAND DECLARATION STARTS HERE====================>)")

	// declaring variables using the Walrus symbol(:=)
	myName := "John"
	phone := 1234567890
	isSick := true
	firstName := &name
	averagePi := 3.14

	// check the types
	fmt.Printf("here are the variable types: %T and value: %v\n", myName, myName)
	fmt.Printf("here are the variable types: %T and value: %v\n", phone, phone)
	fmt.Printf("here are the variable types: %T and value: %v\n", isSick, isSick)
	fmt.Printf("here are the variable types: %T and value: %v\n", firstName, firstName)
	fmt.Printf("here are the variable types: %T and value: %v\n", averagePi, averagePi)

	// type conversion
	// converting int to float64
	floatPhone := float64(phone)
	fmt.Printf("here are the variable types: %T and value: %v\n", floatPhone, floatPhone)

	// converting float64 to int
	intPhone := int(floatPhone)
	fmt.Printf("here are the variable types: %T and value: %v\n", intPhone, intPhone)
}
