// Package main - makes this package an executable prog.
package main

import (
	"fmt"
	"runtime"
)

/*
func main
Go executes the program using the func. Can be considered entry point into the program
*/
func main() {
	fmt.Println(runtime.NumCPU() + 1)
}
