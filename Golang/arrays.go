package main

import "fmt"

func main() {
	fmt.Println()
	var arr [5]int
	var name [5]string
	arr1 := [5]int{1, 2, 3, 4, 5}
	// indexing starts from zero
	arr[0] = 1
	arr2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(name)
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(len(arr))
	sum_ := 0
	for i := 0; i < len(arr1); i++ {
		sum_ += arr1[i]
	}
	fmt.Println(sum_)
	fmt.Println(arr2)
	fmt.Println(arr2[0][1])
}
