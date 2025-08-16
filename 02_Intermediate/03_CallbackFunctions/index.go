package main

import "fmt"


func check(name string, cool func (string)){
	cool(name)
}


func main(){
	check("Lily", func (na string){
		fmt.Println("Hello My name is :", na)
	})
}