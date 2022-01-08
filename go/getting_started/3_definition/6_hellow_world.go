package main

import "fmt"

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("Bazz")
}

func main() {
	bazz()
	fmt.Println("Hello world!", "TEST TEST")
}
