package main

import (
	"fmt"
	f "fmt"
)

// var enabled bool   <-- [enabled redeclared in this block
//	                           previous declaration at ./dup.go:8:5g]

func main() {

	fmt.Println("Hello Gophers!!")
	f.Println("Printing using f instead of fmt")
	dup()
}
