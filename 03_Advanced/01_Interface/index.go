package main

import "fmt"

// Define an interface named Speaker
type Speaker interface {
	Speak() string
}

// Implement the Speaker interface for Dog
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Implement the Speaker interface for Cat
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Function that takes any type implementing the Speaker interface
func announce(speaker Speaker) {
	fmt.Println("The speaker says:", speaker.Speak())
}

func main() {
	// Create instances of Dog and Cat
	dog := Dog{}
	cat := Cat{}

	// Use the announce function with different types
	announce(dog) // The speaker says: Woof!
	announce(cat) // The speaker says: Meow!
}
