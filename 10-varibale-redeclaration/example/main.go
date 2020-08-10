package main

import "fmt"

func main() {
	name := "Akash"
	fmt.Println(name)
	name, age := "Liza", 38
	fmt.Println("name redeclared to", name, "of age", age)

	// name = "Raji"  // value assigned to name
	// YOB := 1977    // new variable reinitialized

	// ^------- same as above
	name, YOB := "Raji", 1977
	fmt.Println("value assigned to name:", name, "whose YOB is:", YOB)
}
