package main

import (
	"fmt"
)

type Person struct {
	name string
	state string
}

func (p Person) Details () string {
	return p.name + " and state is : " + p.state
}

func main(){
	person := Person{
		name: "Brook",
		state: "Black Sea",
	}

	details := person.Details()

	fmt.Println(details)
}