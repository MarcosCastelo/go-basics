package main

import "fmt"

func F1002() {
	PI := 3.14159
	var r float64
	fmt.Scan(&r)
	fmt.Printf("A=%.4f", PI*r*r)
}
