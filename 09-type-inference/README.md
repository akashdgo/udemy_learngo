### Type Inference

- variable initialization 

```
package main
import "fmt"

func main() {
    var safe bool = true <-- variable init
    fmt.Println(safe)
}
```

__Output__

```
ad@mac$go run main.go
true
```
- type inference
```
package main
import "fmt"

func main() {
    var safe = true 
    fmt.Println(safe)
}
```

__Output__

```
ad@mac$go run main.go
true
```
- The __type inference__ type inference means that Go can figure out the type of the variable automatically.

![image](https://user-images.githubusercontent.com/28204484/89301245-13d6e180-d687-11ea-9c36-80d08664b5d7.png)

<-- __short declaration statement__
- short declaration value should always have a value
- short declaration can only be used inside a function and not on the package level

![image](https://user-images.githubusercontent.com/28204484/89303096-55688c00-d689-11ea-856b-6b2546976f87.png)

- __package scope varibale__ will stay in the computer memory till the program is terminated as compared to the __block scope variable__ which is removed from the memory once the block ends.
- reason why too many __package var__ in a program.
- at the package level, all statements start with a keyword.
    - so short declaration cannot be used at package level.

![image](https://user-images.githubusercontent.com/28204484/89303864-606fec00-d68a-11ea-985f-6ecb8c00fb74.png)


