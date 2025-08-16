package main

import "fmt"

func show(){
	fmt.Println("Function called")
}

func sum (a int, b int){
	fmt.Println(a + b)
}

// ! This will take the infinite Number of parameters
func cool(num ...int){
	// * But it comes in array
}

func cool2 (name string, marks ...int){
	// There are two params
	// 1. String 
	// 2. Infinite numbers
}

func multiply (a int, b int) int {
	return a * b
}



func main(){
	func(){
		fmt.Println("IFI function")
	}()
}