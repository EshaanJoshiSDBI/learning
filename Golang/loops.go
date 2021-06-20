package main

import "fmt"

func main() {
	x := 1
	for x < 5 {
		fmt.Println(x)
		x++
	}
	for x := 0; x < 5; x++ {
		fmt.Println(x)
	}
	// break exits the loop
	// continue skips everything below, it jupms to the beginning of the loop
	for x := 0; x <= 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			continue
		}

	}
}
