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

	arr4 := [5]int{}     //not initialized, default value is 0
	arr5 := [5]int{1, 2} //partially initialized

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println("the length is: ", len(arr5))
}
