package main

import "fmt"

func main() {
	// slicers lets us take portions of an array
	// can change the size of slicers
	// var x [5]int = [5]int{1, 2, 3, 4, 5}
	// var s []int = x[1:3] // no values in left and right of : means copy the entire array, no value at left means start at the beginning, no value at right means go till the end.
	// fmt.Println(len(s))
	// fmt.Println(s[:cap(s)])
	var a []int = []int{5, 6, 7, 8, 9}
	a = append(a, 10)
	fmt.Println(cap(a[:3]))
	fmt.Println(a)
	b := make([]int, 5)
	fmt.Println(b)
	fmt.Printf("%T", b)
}
