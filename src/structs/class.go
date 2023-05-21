package main

import (
	"fmt"
)

type Employee struct {
	id int
	name string 
}

func main() {
	e := Employee{}
	e.SetId(1)
	e.SetName("Germán")
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())
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