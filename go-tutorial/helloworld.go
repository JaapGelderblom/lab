package main

import (
	"fmt"
)

func main() {
	/*some long multiline
	comment over multiple lines*/
	var a, b string = "Hello", "World"
	c, d := "another string", "and another" // type is inferred
	fmt.Print("Hello World!")
	fmt.Print(a, b)
	fmt.Println(a, b)
	fmt.Println("something else")
	fmt.Println(c, d)
}
