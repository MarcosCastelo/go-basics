package main

import "fmt"

func F1006() {
	var a float64
	var b float64
	var c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	average := ((a * 2) + (b * 3) + (c * 5)) / 10
	fmt.Printf("MEDIA = %.1f\n", average)
}
