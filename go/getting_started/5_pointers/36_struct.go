package main

import "fmt"

type Vertex struct {
	X, Y int
	S    string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	// (*v).X = 1000
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

	/*
		    s := make([]int, 0)
			s := []int{}
			fmt.Println(s)
	*/

	v8 := Vertex{1, 2, "test"}
	changeVertex(v8)
	fmt.Println(v8)

	v9 := &Vertex{1, 2, "test"}
	changeVertex2(v9)
	fmt.Println(v9)

}
