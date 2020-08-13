package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Redeclare
//
// 	1. Short declare two int variables: age and yourAge
//     (use multiple short declaration syntax)
//
//  2. Short declare a new float variable: ratio
//     And, change the 'age' variable to 43
//
//     (! You should use redeclaration)
//
//  4. Print all the variables
//
// EXPECTED OUTPUT
//  43, 20, 3.14
// ---------------------------------------------------------

func main() {
	// ADD YOUR DECLARATIONS HERE
	age, yourAge := 40, 38
	age, ratio := 43, 0.88

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(age, yourAge, ratio)
}
