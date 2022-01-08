/*
class Vertex(object):

    def __init__(self, x, y):
        self._x = x
        self._y = y

    def area(self):
        return self._x * self._y

    def scale(self, i):
        self._x = self._x * i
        self._y = self._y * i

v = Vertex(3, 4)
v.scale(10)
print(v.area())
*/
package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

//func Area(v Vertex) int {
//	return v.x * v.y
//}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func main() {
	// v := Vertex{3, 4}
	// fmt.Println(Area(v))
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v.Area())
}
