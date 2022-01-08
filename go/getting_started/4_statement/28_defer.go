package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

func main() {
	defer fmt.Println("world")
	foo()
	fmt.Println("hello")

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	file, _ := os.Open("./lesson.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
