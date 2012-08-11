package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		res[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			x, y := float64(i), float64(j)
			res[i][j] = uint8(x * y)
		}
	}
	return res
}

func main() {
	fmt.Println(Pic(16, 16))
}
