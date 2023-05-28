package main

import (
	"fmt"
	"time"
)

func main() {
	//anonymous function
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)
	
	//go rutine
	c := make(chan int)
	go func() {
		fmt.Println("Starting func")
		time.Sleep(5 * time.Second)
		fmt.Println("END func")
		c <- 1
	}()
	<-c

}