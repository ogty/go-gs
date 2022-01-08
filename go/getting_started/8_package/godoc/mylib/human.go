package mylib

import "fmt"

var Public string = "Public"
var private string = "private"

type Person struct {
	// Name
	Name string
	// Age
	Age int
}

func Say() {
	fmt.Println("Human!")
}
