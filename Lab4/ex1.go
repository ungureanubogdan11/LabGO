package main

import (
	"fmt"

	"slices"
)

func statistici(numere []int) (int, int, float32) {
	sum := 0
	for _, x := range numere {
		sum += x
	}
	mini := slices.Min(numere)
	maxi := slices.Max(numere)
	l := len(numere)
	medie := float32(sum) / float32(l)
	return mini, maxi, medie
}

func main() {
	x := []int{4, 8, 15, 16, 23, 42}
	a, b, c := statistici(x)
	fmt.Println(a, b, c)
}
