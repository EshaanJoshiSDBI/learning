package main

import "fmt"

func main() {
	//if condition{}
	age := 19
	//fmt.Println("Before if")
	if age >= 18 {
		fmt.Println("You can ride")
	} else if age >= 14 {
		fmt.Println("You can ride with a parent")
	} else {
		fmt.Println("You cannot ride")
	}

	// if age >= 18 {
	// 	fmt.Println("Can ride")
	// } else {
	// 	fmt.Println("Can't ride")
	// }
}
