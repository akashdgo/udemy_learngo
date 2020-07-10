### Importing

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
        - *"fmt.Println"*
        - *"errors.Printerr"*
        - *"time.PrintTime"*
   calls.

- When a package is imported, its files  will be copied into the final, compiled binary.

### File Scope

** main.go
```
package main
import "fmt"
func main() {
	hey()
	fmt.Println("Hello!!") <--- fmt imported in "package main"
	bye()
}
```

**  bye.go
```
package main
// import "fmt"  <-- fmt not imported & commented out. For bye.go to be able to use it this statement needs to be present.
func bye() {
	fmt.Println("bye!!!")
}
```

- the above will thrown an error since fmt is not imported 
- in main.go fmt is imported but has a file scope and has the perimeter of file main.go and cannot be called from hey.go even though it is part of pkg main.

![image](https://user-images.githubusercontent.com/28204484/87104512-ccbb2380-c275-11ea-9ef2-a6e070b34b08.png)