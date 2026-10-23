package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	var target int
	fmt.Scan(&target)

	m := map[int]int{}
	for _, x := range arr {
		m[x] = 1
	}

	for x := range arr {
		if m[target-x] > 0 {
			fmt.Println("am gasit nr")
			fmt.Println(x, target-x)
		}
	}

}
