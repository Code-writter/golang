# Functions

A function is a block of statements that can be used repeatedly in a program.

Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.

```go
package main
import "fmt"

// Function Defination
func printName(username string) {
    fmt.Println(username)
}

func main() {
    // Function call
    printName("HuXn")
}
```

## Return 

 > func function_name (params) return_type {
 >   return _ value
 > }

```go
func multiply (a int, b int) int {
	return a * b
}
```

# Annonymus Function

```go
func main(){
	anotherFunc := func(a int, b int) int {
		return a - b
	}

	ans := anotherFunc(3, 5)
	fmt.Println(ans)
}
```

# IFI
```go
	func(){
		fmt.Println("IFI function")
	}()
```