package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU() + 1)
}

// Output

// akashdas@akash-mac:02-example$ go run main.go
// 4
// akashdas@akash-mac:02-example$ go run main.go
// 5
//akashdas@akash-mac:02-example$
