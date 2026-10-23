package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("cuvinte.txt")
	text := string(data)
	cuvinte := strings.Fields(text)
	fmt.Println(cuvinte)

	m := map[string]int{}
	for _, c := range cuvinte {
		m[c] = m[c] + 1
	}

	fmt.Println(m)

}
