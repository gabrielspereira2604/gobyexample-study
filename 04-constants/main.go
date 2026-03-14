package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// n = 123 // This would cause a compile error since constants cannot be reassigned
	// s = "other" // This would also cause a compile error since constants cannot be reassigned

	// 3e20 é notação científica → 3 x 10^{20}
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
