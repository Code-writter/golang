package main

import "fmt"

func sendDataInChannel(data chan<- int) {
	// Step 2 : Put data in the channel
	data <- 20
}

func main() {
	// // Step 1 : Create channel
	// channel := make(chan int)

	// // fmt.Println(channel) // Its basically a reference address in memory for commuincation
	// // by shared memory

	// // pass in Go routine
	// go sendDataInChannel(channel)

	// // Step 3 : Recieve data from the channel
	// valueWeReciveFromChannel := <-channel

	// fmt.Printf("Value get is : %d", valueWeReciveFromChannel)

	unbufferedChannel := make(chan int)

	bufferedChannel := make(chan int, 2)

	go func() {
		unbufferedChannel <- 100
		close(unbufferedChannel)
	}()

	go func() {
		bufferedChannel <- 10
		bufferedChannel <- 20
		// we can close the channel
		close(bufferedChannel)
	}()
	
	for data := range bufferedChannel{
		fmt.Println(data)
	}

}	