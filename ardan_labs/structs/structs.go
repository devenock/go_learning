package structs

import "fmt"

func main() {
	type Car struct {
		color string
		model string
		year  int
		isNew bool
	}
	fmt.Printf("The car is of type %T\n", Car{})
}
