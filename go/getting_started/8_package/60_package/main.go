package main

import (
	"fmt"

	"awesomeProject/mylib"
	"awesomeProject/mylib/under"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	under.Hello()
}
