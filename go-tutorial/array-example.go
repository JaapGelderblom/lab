package main

import (
	"fmt"
)

func main() {
	//examples of arrays:
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	//inferred length:
	var arr3 = [...]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
