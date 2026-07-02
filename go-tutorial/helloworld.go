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

	fmt.Printf("The type of c is: %T and the value is: %v\n", c, c)

	var i = 15

	fmt.Printf("%b\n", i)  //binair
	fmt.Printf("%d\n", i)  //decimal
	fmt.Printf("%+d\n", i) //decimal with sign
}
