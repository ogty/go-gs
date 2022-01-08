package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
}
