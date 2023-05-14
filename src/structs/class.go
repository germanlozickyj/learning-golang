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
	e.id = 0
	e.name = "German"
	fmt.Println(e)
}