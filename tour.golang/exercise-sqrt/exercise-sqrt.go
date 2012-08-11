package main

import (
	"fmt"
	"math"
)

// first attempt
func Sqrt_iter(x float64, iters int) float64 {
	res := 1.0
	for i := 0; i < iters; i++ {
		res = res - (res*res-x)/(2*res)
	}
	return res
}

func Sqrt(x float64) (int, float64) {
	res := 1.0
	i := 0
	for delta := 1.0; delta > .00001; i++ {
		new_res := res - (res*res-x)/(2*res)
		delta = math.Abs(res - new_res)
		res = new_res
	}
	return i, res
}

func main() {
	for i := 0; i < 100; i++ {
		x := float64(i)
		iters, res := Sqrt(x)
		fmt.Printf("sqrt(%f) == %f [%d]\n",
			x, res, iters)
	}
}
