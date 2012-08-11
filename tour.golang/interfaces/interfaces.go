package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	f := MyFloat(4)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v // *Vertex implements Abser
	fmt.Println(a.Abs())

	//a = v // Vertex does not
}
