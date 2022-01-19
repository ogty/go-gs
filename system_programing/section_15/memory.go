package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Page Size: %d\n", os.Getpagesize())
}
