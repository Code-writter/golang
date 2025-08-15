package main

import (
	"fmt"
)

func main(){

	var arr1 = [10]int{20, 11, 15}
	 
	// arr2 := []int {10, 20, 30}

	// arr3 := []int{}
	
	fmt.Println(arr1)
	arr1[3] = 100
	fmt.Println(arr1)
	fmt.Println(len(arr1))

	for i := 0; i<len(arr1); i++ {
		fmt.Println(arr1[i])
	}
}
