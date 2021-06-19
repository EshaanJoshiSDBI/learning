package main

import "fmt"

func main() {
	// fmt.Printf("Hello %t %b %x %e %q\n", true, 10, 16, 2.718, "Hey")
	var out string = fmt.Sprintf("Hello %t %b %x %e %q\n", true, 10, 16, 2.718, "Hey")
	fmt.Println(out)
}
