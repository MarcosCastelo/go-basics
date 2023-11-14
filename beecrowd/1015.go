package main

import (
	"fmt"
)

func F1015() {
	var a int64
	var b float64

	fmt.Scan(&a, &b)

	result := float64(a) / b
	fmt.Printf("%.3f km/l \n", result)
}
