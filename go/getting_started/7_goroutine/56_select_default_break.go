package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
			// return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("##############")
}
