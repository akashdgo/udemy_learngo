### Go Docs
-----------
- always start the comment or documentation with the name of the declartion
- this will help in creating the documentation automatically

```
// **_Package main_** - makes this package an executable prog.
package main

import (
	"fmt"
	"runtime"
)

/*
**_func main_**
Go executes the program using the func. Can be considered entry point into the program
*/
func main() {
	fmt.Println(runtime.NumCPU() + 1)
}
```

To view the docuemtation of NumCPU func.

```
ad@mac:02-example$ `go doc -src runtime NumCPU`
package runtime // import "runtime"

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
func NumCPU() int {
        return int(ncpu)
}
```Similarly to view the documentation of fmt package

```
ad@mac:02-example$ go doc -src fmt Println
package fmt // import "fmt"

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
        return Fprintln(os.Stdout, a...)
}
```