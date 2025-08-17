package main

import "fmt"

func main() {
    // Declare a variable
    x := 42

    // Create a pointer that stores the memory address of x
    var pointerToX *int = &x

    // Print the value and memory address of x
    fmt.Println("Value of x:", x)
    fmt.Println("Memory address of x:", &x)

    // Print the value and memory address stored in the pointer
    fmt.Println("Value pointed to by pointerToX:", *pointerToX)
	// This gives the value of the x
    fmt.Println("Memory address stored in pointerToX:", pointerToX)
}