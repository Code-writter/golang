# String Data Type

### The string data type is used to store a sequence of characters (text).

 > String values must be surrounded by double quotes

- single character can be accessed by the zero base indexing 
```go
package main
import "fmt"

func main(){
    name := "Abhishek"
    fmt.Println(name[0]) // This will print Uni Code value of A
    fmt.Printf( "%c",name[0]) // This will print A
}
```

```go

package main
import "fmt"

func main() {
    fmt.Println("I'm a String data type")
    info := "Hello everyone"
    fmt.Println(info)
}
```

### For mutlti line strings use the back-ticks sign (``)

```go
package main
import "fmt"


func main(){
        
    name := `This is a example of the 
    multi 
    line string`
    fmt.Println(name) // This will print 
}

```

## We can use the (+) operator to concatinate two or more strings
```go
package main
import "fmt"

func main(){
    msg1 := "this is string one"
    msg2 := "this is string two"

    fmt.Prinln(msg1 + msg2)

}
```

## Compare two strings are identical or not

```go
package main
import(
    "fmt"
    "string"
)
func main(){
    msg1 := "hello"
    msg2 := "hello"

    fmt.Println(string.Compare(msg1, msg2)) // This returns the boolean value
}
```
## Search the sub-string in the String

```go
package main
import(
    "fmt"
    "string"
)
func main(){
    msg1 := "This is great"

    fmt.Println(string.Contains(msg1, "This")) // This returns the boolean value
}
```
