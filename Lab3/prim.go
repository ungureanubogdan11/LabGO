package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	ok := true

	for d := 2; d < x; d++ {
		if x%d == 0 {
			ok = false
		}
	}

	if ok {
		fmt.Println("Numarul este prim")
	} else {
		fmt.Println("Numarul nu este prim")
	}

}
