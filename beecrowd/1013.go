package main

import (
	"fmt"
	"math"
)

func greater(a int64, b int64) int64 {
	return (a + b + int64(math.Abs(float64(a-b)))) / 2
}

func F1013() {
	var a, b, c int64

	fmt.Scan(&a, &b, &c)

	result := greater(greater(a, b), c)
	fmt.Printf("%d eh o maior\n", result)
}
