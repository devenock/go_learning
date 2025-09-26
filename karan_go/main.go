package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello Karan Go!")

	// Go was developed to basically bring a blend of static typing and dynamic programming.
	// The google team wanted to combine the best parts of python and c++ so that they could
	// build reliable systems that took advantage of milti-core processprs.

	// you can always initialize a go module using: go mod init <module name>.
	// A module can be basically defined as a collection of Go packages.

	// VARIABLE AND DATA TYPES
	// We mainly have 4 primitive data types in Go: int, string, bool, float
	// byte => uint8
	// rune => int32(represent a unicode code point)
	// float32
	// float64(default floating point value)
	//
	// ZERO VALUES
	var name string
	var age int
	var isMarried bool
	var averageWeight float64

	// print values
	// fmt.Printf("Name: %v\n", name)
	// fmt.Printf("Age: %v\n", age)
	// fmt.Printf("IsMarried: %v\n", isMarried)
	// fmt.Printf("AverageWeight: %v\n", averageWeight)
	//
	// OR
	fmt.Printf("%+v\n", name)
	fmt.Printf("%+v\n", age)
	fmt.Printf("%+v\n", isMarried)
	fmt.Printf("%+v\n", averageWeight)

	// TYPE CONVERSION
	// float64 to int
	newWeight := int(averageWeight)
	fmt.Println(newWeight)

	// int to float64
	newAge := float64(age)
	fmt.Println(newAge)

	// int to string
	strAge := strconv.Itoa(age)
	fmt.Println(strAge)

	// string to int
	intAge, _ := strconv.Atoi(strAge)
	fmt.Println(intAge)

	// TYPE ALIAS
	type MyAlias = string
	var str MyAlias = "My name is noki"
	fmt.Println(str)

	// DEFINED TYPES
	type MyString string
	var name1 MyString = "This is my string"
	fmt.Println(name1)

	// the major difference between type aliases and defined types is because
	// defined types does not use the "=" sign.

	// FLOW CONTROL
	//
	// if/else

	// x := 10
	// if x > 5 {
	// 	fmt.Println("The number x is greater than 5")
	// }
	//
	// OR COMPACT IF
	if x := 10; x < 20 {
		fmt.Println("X is less than 20")
	}

	// switch

	// day := "Friday"

	// switch day {
	// case "Monday":
	// 	fmt.Println("It is the first day of the week")
	// case "Tuesday":
	// 	fmt.Println("It is the second day of the week")
	// default:
	// 	fmt.Println("No rest for us today!")
	// }
	//
	// SWITCH WITH SHORTHAND
	switch day := "Wednesday"; day {
	case "Wednesday":
		fmt.Println("It is the third day of the week")
	case "Friday":
		fmt.Println("It is the sixth day of the week")
	default:
		fmt.Println("That day does not exist!")
	}

	// SWITCH WITH NO CONDITION
	day := "Wednesday"
	switch {
	case day == "Wednesday":
		fmt.Println("It is the third day of the week")
	case day == "Friday":
		fmt.Println("It is the sixth day of the week")
	default:
		fmt.Println("That day does not exist!")
	}

	// LOOPS
	// In Go we only have one type of loop which is the for loop
	// It is very versatile and same as the if statement, fo loop does not need parenthesis ()
	// the basic for loop has got three components separated by semicolons:
	// 1. init statement: executed before the first iteration
	// 2. condition expression: evaluated before every iteration
	// 3. post statement: executed at the end of every iteration
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// BREAK AND CONTINUE
	for i := 0; i < 20; i++ {
		if i > 5 {
			continue
		}
		fmt.Println(i)

		if i > 15 {
			break
		}
	}
	fmt.Println("The loop was broken!!")

	// continue is used when we want to skip the remaining portion of the loop and break is used
	// when we want to break out of the loop.

	// init and post statements are optional
	i := 9
	for i < 5 {
		i += 3
	}

	// forever loop
	// If we omit the loop confition, it loops forever. This is known as forever loop
	//
	for {
		// do stuff
	}

	// FUNCTIONS
}
