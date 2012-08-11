package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 9, 11}
	fmt.Println(p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println(p[1:3])
	fmt.Println(p[2:])
	fmt.Println(p[:2])
	fmt.Println(p, len(p), cap(p))
}
