/*
<
>
<=
>=
==
!= */
package main

import "fmt"

func main() {
	x := 5
	y := 6
	val := x <= 4
	val2 := y >= 3
	fmt.Println(val)
	fmt.Println(val2)
	a := "a"
	b := "b"
	fmt.Printf("%t\n", a < b)
}
