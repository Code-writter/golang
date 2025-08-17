package main

import "fmt"


func printNumber(a int, b int) (int, int){
	return a, b
}
func printString(a string, b string) (string, string){
	return a, b
}
func printBool(a bool, b bool) (bool, bool){
	return a, b
}


func printAll[T any] (a T, b T) (T, T){
	return a, b
}


func main() {
	// num1, num2 := printNumber(4, 5)
	// string1, string2 := printString("Monkey D", "Luffy")
	// bool1, bool2 := printBool(true, true)

	num1 , num2 := printAll(4, 50)
	// num1 , num2 := printAll[int](4, 50)

	// string1, string2 := printAll("Ronoroa", "Zoro")
	string1, string2 := printAll[string]("Ronoroa", "Zoro")

	bool1, bool2 := printAll(true, true)

	fmt.Println(num1, num2)
	fmt.Println(string1, string2)
	fmt.Println(bool1, bool2)
}