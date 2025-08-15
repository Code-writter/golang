package main

import "fmt"

func main(){
	users := []int {10, 20, 30, 40, 50}

	fmt.Println(users) //[10 20 30 40 50]
	fmt.Println(users[:]) //[10 20 30 40 50]
	fmt.Println(users[1 : 3]) //[20 30] (skiped the 0th and start form 1, skiped the rest formt the 3)

	fmt.Println(users[2:]) // [30 40 50]

	// * insert Element at the end of the slice
	users = append(users, 660)
	fmt.Println(users) //[10 20 30 40 50 660]

	test := []int{}
	test = append(test, 40, 50, 60)

	// fmt.Println(users)
	fmt.Println(test)

	check := make([]string, 2, 3)

	fmt.Println(check)
	check[0] = "Hello"
	check[1] = "bye"
	fmt.Println(check)


	// * Way to iterate slice
	// for i := 0; i<len(users); i++ {
	// 	fmt.Println(users[i])
	// }

	for _ , value := range users {
		fmt.Println(value)
	}
}