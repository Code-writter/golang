package main

import "fmt"

type Views struct{
	num int
}

func (v *Views) inc(){
	v.num += 1
}



func main(){
	test := Views{num: 0}
	test.inc()
	fmt.Println(test.num)
}