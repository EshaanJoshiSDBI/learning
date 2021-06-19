package main

import (
	"fmt"
	// for more maths stuff
	"math"
)

func main() {
	var num1 int = 8
	var num2 int = 4
	// all operands must be same type
	answer := num1 + num2
	fmt.Printf("%d\n", answer)
	var num3 float64 = 8
	var num4 int = 4
	answer2 := num3 / float64(num4)
	fmt.Println(answer2)
	var num5 float32 = 9
	var num6 float32 = 4
	// as num5 and num6 are int num5/num6 returns 2
	fmt.Println(num5 / num6)
	fmt.Println(num1 % num2)
	fmt.Println(math.Cos(3.1415))
}
