package main

import (
	"fmt"
)

func main() {
	arr := []float32{1, 2, 3, 4, 5, 6}

	var index int
	fmt.Scan(&index)

	ans := append(arr[:index], arr[index+1:]...)

	fmt.Println(ans)

}
