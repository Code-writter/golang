package main

import "fmt"

func main(){
	num := 5

	if num > 5 {
		fmt.Println("Number is grater than 5")
	}else if num < 5 {
		fmt.Println("Number is less than 5")
	}else{
		fmt.Println("Number is equal to 5")
	}

	switch num {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Other Number")
	}

	
}