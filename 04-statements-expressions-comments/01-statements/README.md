### Statements
--------------

- Statements are like instructions they tell Go to execute something
- Actually they control the execution flow in a Go program.
- It means that statements can start, stop, repeat the execution of code. 
- In every code line there should be only one statement at a time.
- Some statements like an 'if statement' for example can have a block of code or a body.


<-- every line is a statement 
<-- "Hello Gophers!!!" is part of the statement (expressions)
```
package main
import "fmt"

// statement block and it starts with { and ends with } 
func main() {                         
	fmt.Println("Hello Gophers!!!")  
// another block [sub-block]
    if 5 > 3 {
        fmt.Println("bigger")
    }   
}                                         
```

<-- all of the above are all **_declaration statement_** since they do not change the execution flow of the code.
<-- they just declare new names. 
<-- 'Println' call changes executions since it goes into Println func and executes it code.


![image](https://user-images.githubusercontent.com/28204484/87133780-72da4e00-c2b5-11ea-9742-5e5d3f70d85e.png)

![image](https://user-images.githubusercontent.com/28204484/87134101-d2d0f480-c2b5-11ea-95b6-47e0e02fc9f3.png)


```
package main

import "fmt"

func main() {
	fmt.Println("hello"); fmt.Println("Akash!!!")
}
```
when saved to be run is formatted to

```
package main

import "fmt"

func main() {
	fmt.Println("hello")
	fmt.Println("Akash!!!")
}
```

**Output**
```
ad@mac:ch-prog$ go run main.go
hello
Akash!!!
```
