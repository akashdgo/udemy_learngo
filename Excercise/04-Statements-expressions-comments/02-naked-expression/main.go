// ---------------------------------------------------------
// EXERCISE: Naked Expression
//
//  1. Try to type just "Hello" on a line.
//  2. Do not use Println
//  3. Observe the error
//
// ---------------------------------------------------------

package main

func main() {
	// "Hello"
}

/*
Output
error: "Hello" evaluated but not used go
*/

// Because:
// "Hello" literal returns a value but there isn't any
// statement which uses it.
//
// So:
// You can't use expressions alone without statements.
