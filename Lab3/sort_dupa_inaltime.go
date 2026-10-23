package main

import (
	"fmt"
	"sort"
)

func main() {

	m := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	M := map[int]string{}
	for i, j := range m {
		M[j] = i
	}

	arr := []int{}

	for _, i := range m {
		arr = append(arr, i)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	for _, x := range arr {
		fmt.Println(M[x], x)
	}

}
