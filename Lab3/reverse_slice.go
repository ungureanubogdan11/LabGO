package main

import (
	"fmt"
	"strings"
)

func main() {
	S := "  Racecar    "
	s := strings.TrimSpace(S)
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			fmt.Println("nu este palindrom")
			return
		}
	}

	fmt.Println("este palindrom")

}
