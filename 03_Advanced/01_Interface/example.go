package main

import "fmt"

type University interface {
	course() string
}


// college one

type SAIT struct{}

func (s SAIT) course() string {
	return "B-tech"
}

// another college

type MIST struct{}

func (m MIST) course() string {
	return "Something"
}


func main(){

	s := SAIT{}
	m := MIST{}

	fmt.Println(s.course())
	fmt.Println(m.course())
}