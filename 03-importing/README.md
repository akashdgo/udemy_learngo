### Importing#1 
---------------

1. Importing only happens at the file level and not at pkg level.
2. Importing == declaring what is inside the improted package files in your own file.

```
import {
    "fmt"
    "errors"
    "time"
}
```

- exported names as if they're your own e.g. 
    - **"fmt.Println"**
	- **"errors.Printerr"**
	- **"time.PrintTime"**
   calls.

- When a package is imported, its files  will be copied into the final, compiled binary.

### File Scope

**main.go**
```
package main
import "fmt"
func main() {
	hey()
	fmt.Println("Hello!!") <--- fmt imported in "package main"
	bye()
}
```

**bye.go**
```
package main
// import "fmt"  <-- fmt not imported & commented out. For bye.go to be able to use it this statement needs to be present.
func bye() {
	fmt.Println("bye!!!")
}
```

- the above will thrown an error since fmt is not imported 
- in main.go fmt is imported but has a file scope and has the perimeter of file main.go and cannot be called from hey.go even though it is part of pkg main.

![image](https://user-images.githubusercontent.com/28204484/87104745-7a2e3700-c276-11ea-850d-e556b18452a6.png)

### Importing#2
----------------

- imports in a Go file requires unique name and same package with the same name cannot be imported more than once.

```
package main
import "fmt"
import "fmt" <---- cannot have the same name
```

- to import it again - it is required to rename it as follows:

```
package main
import "fmt"
import f "fmt" <---- go file will import "fmt" as 'f': so "fmt" pkg can be used by 'f' or 'fmt'
```

```
package main

import (
	"fmt"
	f "fmt"
)

func main() {

	fmt.Println("Hello Gophers!!")
	f.Println("Printing using 'f' instead of 'fmt'")
}
```
**_Output_**

```
ad@mac:03-importing$ go run main.go
Hello Gophers!!
Printing using 'f' instead of 'fmt'
```
