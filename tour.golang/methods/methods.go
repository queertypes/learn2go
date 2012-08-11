package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// using a pointer receiver, the method efficiently
// receives its argument and it can be modified
func (v *Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

// using a value receiver, the method copies in
// its argument and cannot modify it
func (v Vertex) Scale_(f float64) {
	v.X *= f
	v.Y *= f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Abs_() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

	vv := Vertex{3, 4}
	vv.Scale_(5)
	fmt.Println(vv, vv.Abs_())
}
