package main

import (
	"fmt"
	"time"
)


func printCount(){
	for i := 0; i<10; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}


func main(){
	printCount()

	// To make it Go-routine
	go printCount()
}