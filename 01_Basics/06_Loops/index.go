package main

import "fmt"

func main(){
	for i := 0; i<10; i++ {
		fmt.Println(":::: >")
	}

	// There is now while loop so
	num := 0
	for num < 10 {
		if num == 5 {
			// * we stop loop when num == 5 
			break
		}
		fmt.Println(num)
		num++
	}
}