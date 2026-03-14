package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty:", s, len(s) == 0, "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	l := s[2:5]
	fmt.Println("slice:", l)

	// slice de início até o índice 5
	l = s[:5]
	fmt.Println("slice:", l)

	// slice do índice 2 até o final
	l = s[2:]
	fmt.Println("slice:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
