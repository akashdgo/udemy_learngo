### Comments
------------

- to provide comments and documentation for code.
- // or for block of comments /* <snippet> */

```
package main
// imports package below into 'package main' 
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU() + 1)
}

/* 
Output

ad@mac:02-example$ go run main.go
4
ad@mac:02-example$ go run main.go
5
ad@mac:02-example$

*/
```