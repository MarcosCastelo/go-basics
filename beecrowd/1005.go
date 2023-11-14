package main

import "fmt"

func F1005() {
	var a float64
	var b float64
	fmt.Scan(&a)
	fmt.Scan(&b)

	average := ((a * 3.5) + (b * 7.5)) / 11
	fmt.Printf("MEDIA = %.5f\n", average)
}
