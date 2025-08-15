# While Loop

> Unlike other programming languages, Go doesn't have a dedicated keyword for a while loop. However, we can use the for loop to perform the functionality of a while loop.

- for loop
```go
	for i := 0; i<10; i++ {
		fmt.Println(":::: >")
	}
```

- while loop
```go
// Program to print numbers between 0 and 10
package main
import ("fmt")

func main() {
  number := 0

  for number <= 10 {
    fmt.Println(number)
    number++
  }
}
```

# Continue and Break

- Continue keyword is used to skip the specific condition in the looping condition

```go
	num := 0
	for num < 10 {
		if num == 5 {
			// * we skip the 5 
			num++
			continue
		}
		fmt.Println(num)
		num++
	}

```

- Break keyword is used to break through the loop (stop the loop and move forward)

```go
	num := 0
	for num < 10 {
		if num == 5 {
			// * we stop loop when num == 5 
			break
		}
		fmt.Println(num)
		num++
	}

```