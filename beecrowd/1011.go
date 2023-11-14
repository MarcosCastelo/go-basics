package main

import "fmt"

func F1011() {
	var r uint64
	fmt.Scan(&r)

	vol := (4.0 * 3.14159 * float64(r*r*r)) / 3
	fmt.Printf("VOLUME = %.3f\n", vol)
}
