package basics

import "fmt"
import "strconv"

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println("variables in golang", x)
	fmt.Println("variables in golang", y)
	//in golang there aren't try-catch 
	//in golang we drive like this the errors:
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Println("---error----", err)
	} else {
		fmt.Println(myValue)
	}
	//map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println("---map----", m["Key"])
	//loops with range in golang
	s := []int{1,2,3}
	for index, value := range s {
		fmt.Println(index, value)
	}
	//adding value to slice
	s = append(s, 5)
	fmt.Println(s)
}