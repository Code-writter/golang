
package main

import (
	"fmt"
)

// There are two ways to decleare variables in go
// var variable_name = value
// variable_name := value

func main(){
	var name = "Hello"
	// type inferencing
	var age int = 30
	var isStudent bool = true
	// := punisher symbol
	surname := "bye"

	// this cannot be done
	// test string := "gooooo" 

	// * multiple variables

	var(
		first, second , third, fourth = 1, 2, 3, 4
	)
	// var(
	// 	id int = 1
	// 	user string = "hello world"
	// 	isGood bool = true
	// )

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isStudent)
	fmt.Println(surname)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
	fmt.Println(fourth)
}