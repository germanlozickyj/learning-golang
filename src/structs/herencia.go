package main

import "fmt"

type Person struct {
	name string
	age int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.age = 18
	ftEmployee.name = "German"
	ftEmployee.id = 5
	fmt.Println(ftEmployee)
}