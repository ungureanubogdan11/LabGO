package main

import "fmt"

func main() {
	m := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	ans := map[int][]string{}

	for key, value := range m {
		ans[value] = append(ans[value], key)
	}

	fmt.Println(ans)

}
