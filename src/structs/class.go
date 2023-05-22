package main

import (
	"fmt"
)

type Employee struct {
	id int
	name string 
}

func NewEmployee(id int, name string) *Employee {
	return &Employee{
		id : id,
		name : name,
	}
}

func main() {
	e := Employee{}
	e.SetId(1)
	e.SetName("Germán")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
	employee := Employee{
		id : 2,
		name: "Germán 2",
	}
	fmt.Println(employee)
	e4 := NewEmployee(3, "German 3")
	fmt.Println(*e4)
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}