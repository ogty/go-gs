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


class Vertex3D(Vertex):

    def __init__(self, x, y, z):
        super().__init__(x, y)
        self._z = z

    def area_3d(self):
        return self._x * self._y * self._z

    def scale_3d(self, i):
        self._x = self._x * i
        self._y = self._y * i
        self._z = self._z * i


v = Vertex3D(3, 4, 5)
v.scale(10)
print(v.area())
print(v.area_3d())
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

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale3D(10)
	//fmt.Println(v.Area())
	fmt.Println(v.Area3D())
}
