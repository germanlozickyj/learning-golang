package main

import (
	"fmt"
)

//Nos permite crear funciones como slices los argumentos de funciones lo cual no sabemos la longitud exacta
func sum(values ...int) int  {
	total := 0
	for _, num:= range values {
		total += num
	}
	return total
}

//Los retornos con nombre nos permitiran definir variables antes de definir el cuerpo de la funcion, por lo cual solo utilizamos el return para devolverlos
func getValues(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x 
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 5, 10))
	fmt.Println(getValues(2))
}