package main

// Pointer
func employee(name *string, age int){
  if age > 65{
    panic("Age cannot be greater than retirement age")
  }

}

func main() {
  empName := "Samia"
  age := 75
	//  reference variable
  employee(&empName, age)
}