package main

import "fmt"

func main(){
	users := map[string]int {
		"User1" : 100,
	}

	fmt.Println(users) //map[User1:100]
	fmt.Println(users["User1"]) //100

	users["new"] = 100

	fmt.Println(users) //map[User1:100 new:100]

	users["new"] = 500
	fmt.Println(users) //map[User1:100 new:500]

	for key, value := range users {
		fmt.Println(key, value)
	}
}