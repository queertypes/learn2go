package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var res complex128 = 1.0

	for i := 0; i < 100; i++ {
		res3 := cmplx.Pow(res, 3)
		res2 := cmplx.Pow(res, 2)
		res = res - (res3-x)/(3*res2)
	}
	return res
}

func main() {
	fmt.Println(Cbrt(8))
}
