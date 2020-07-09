package main

import "fmt"

func main() {
	fmt.Println("My name is Akash Das")
	fmt.Println("My best friends name is Kumar Muthuswamy")
	fmt.Print("with Print instead of Println\n")

	fmt.Println("akash", "kumar")
	fmt.Print("Using Print instead of Println:\t", "akash", "kumar")
	// fmt.Println("Akash without quotes", akash, "kumar")
	fmt.Println("\n With the pkg main and import at the bottom")
	fmt.Println("error: main.go:3:1: expected 'package', found 'func'")
}
