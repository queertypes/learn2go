package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	strs := strings.Fields(s)
	res := make(map[string]int)

	for _, str := range strs {
		res[strings.ToLower(str)]++
	}
	return res
}

func main() {
	fmt.Println(WordCount("Word count test count Word"))
	fmt.Println(WordCount("CAT CAt CaT Cat cAT cAt caT cat"))
}
