package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var res complex128 = 1.0

	for delta := 1.0; delta > .000001; {
		old_res := res
		res3 := cmplx.Pow(res, 3)
		res2 := cmplx.Pow(res, 2)
		res = res - (res3-x)/(3*res2)
		delta = cmplx.Abs(res - old_res)
	}
	return res
}

func main() {
	fmt.Println(Cbrt(27))
}
