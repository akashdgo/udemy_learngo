## Declaring a variable

- Variable declaration is static and cannot be changed later in the program

```
package main

import (
	"fmt"
)

func main() {
    var speed int
	fmt.Println(speed)
}
```
- variable speed type is int
- every varibale must have a type
- go automatiocally assigns a value to the variable; in the above example 'speed' is assigned a value of '0'

#### output
```
ad@mac:example$ go run main.go
0
ad@mac:example$
``` 
- order of declaration is very important

```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(speed)

    var speed int
}
```

#### output
```
ad@mac:example$ go run main.go
# command-line-arguments
./main.go:9:14: undefined: speed  <-- NOTE
```


