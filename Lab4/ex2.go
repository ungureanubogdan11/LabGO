package main

import "fmt"

func interschimba(a, b *int) {
	x := *a
	*a = *b
	*b = x
}

func dubleaza(a *int) int {
	return 2 * (*a)
}

func main() {
	a, b := 2, 3
	fmt.Println(a, b)
	interschimba(&a, &b)
	fmt.Println(a, b)
	dubleaza(&a)
	fmt.Println(a)
}
