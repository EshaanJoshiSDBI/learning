// AND OR NOT
// && || !
package main

import "fmt"

func main() {
	val := 5 < 7 || 7 > 9 && 8 < 7
	val2 := (5 < 7 || 7 > 9) && 8 < 7
	fmt.Printf("%t \t %t \t %t \n", val, val2, ((true || false) && !false))
}
