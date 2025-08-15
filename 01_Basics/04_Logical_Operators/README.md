# Logical Operators

## Logical operators are used to determine the logic between variables or values.

## 1. Logical and (&&)

| a | b |AND|
| - | - | - |
| 1 | 1 | 1 |
| 1 | 0 | 0 |
| 0 | 1 | 0 |
| 0 | 0 | 0 |

### Returns true if both statements are true

## 2. Logical or (||)

| a | b | OR |
| - | - | - |
| 1 | 1 | 1 |
| 1 | 0 | 1 |
| 0 | 1 | 1 |
| 0 | 0 | 0 |

### Returns true if one statements is true

## 3. Logical not (!)
| a |NOT|
| - | - |
| 1 | 0 |
| 0 | 1 |


### Reverse the result, returns false if the result is true

## 4. Logical XOR (^)
 > Give zero for the similar value

| a | b |XOR|
| - | - | - |
| 1 | 1 | 0 |
| 1 | 0 | 1 |
| 0 | 1 | 1 |
| 0 | 0 | 0 |


```go

package main

import "fmt"

func main() {
	fmt.Println(true && true) // true
	fmt.Println(true && false) // false

	fmt.Println(true || true) // true
	fmt.Println(false || false) // false

	fmt.Println(!true) // false
	fmt.Println(!false) // true
}
```