package main

import "fmt"

func main() {
	ans := -1
	switch ans {
	case 1, -1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("wrong value entered")
	}
	// switch{
	// case ans > 0:
	// 	fmt.Println("Greater than zero")

	// }
}
