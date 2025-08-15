package main

import (
	"fmt"
)

func main(){
	fmt.Println(1 ^ 1) // XOR

	isAdmin := true
	isLoggedIn := true

	// if both is true then this if condition will get fired
	if(isAdmin && isLoggedIn){
		fmt.Println("Welcome Admin")
	}
}