package main

import (
	"fmt"
	f "fmt"
)

var enabled bool

// var enabled bool  <-- [enabled redeclared in this block
//                       	previous declaration at ./dup.go:8:5go]

func dup() {
	enabled = true
	fmt.Println("Hello Gophers!! declaring same variable globally - checking file scope")
	f.Println(enabled)
}
