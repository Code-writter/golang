package main

import (
	"fmt"
)

type Human struct {
	name string
	age int
}

type Student struct {
	Human
	isStudent bool
}


func main(){

	s := Student {
		Human: Human{
			name: "Nami",
			age: 20,
		},
	}

	fmt.Println(s)
	fmt.Println(s.name)

}