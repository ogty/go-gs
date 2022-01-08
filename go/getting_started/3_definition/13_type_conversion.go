package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)
	/*
	_を使うことでエラーを代入する。　第二引数がないとそれだけでerrorになってしまうので、意味の持たない
	_を使い、エラーを回避する。
	おそらく、使うことはないので、普段から_を入れておけばよいだろう...
	*/

	h := "Hello World"
	fmt.Println(string(h[0]))
}
