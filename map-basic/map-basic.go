package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.7, 42.1}
	fmt.Println(m, m["Bell Labs"])

	m["Bell Labs"] = Vertex{41.1, 42.7}
	fmt.Println(m)

	v, ok := m["Bell Labs"]
	fmt.Println("The value:", v, "Present?", ok)

	delete(m, "Bell Labs")
	fmt.Println(m)

	v, ok = m["Bell Labs"]
	fmt.Println("The value:", v, "Present?", ok)

	v, ok = m["Bell L"]
	fmt.Println("The value:", v, "Present?", ok)
}
