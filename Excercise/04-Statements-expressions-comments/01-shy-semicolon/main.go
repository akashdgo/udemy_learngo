// ---------------------------------------------------------
// EXERCISE: Shy Semicolons
//
//  1. Try to type your statements by separating them using
//     semicolons
//
//  2. Observe how Go fixes them for you
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("how are you?")

}

/*
   fmt.Println("Hello"); fmt.Println("how are you?")
*/

// fixes the above to

/*
   fmt.Println("Hello")
   fmt.Println("how are you?")
*/
