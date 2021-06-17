package main

import "fmt"

func main() {
	//implicit
	var number = 260
	number = number + 5
	//printf supports formatting `print formatting`
	fmt.Printf("%T", number)
	//explicit
	var number2 uint32 = 3200
	number2 = number2 + 10
	fmt.Println("\n", "number 2 is", number2)
	//walrus operator
	number3 := 32
	//equivalent to var number3 uint16
	fmt.Println(number3)
	var default_ uint64
	var default_bool bool
	fmt.Println("Default values", "int", default_, "boolean", default_bool)
}
