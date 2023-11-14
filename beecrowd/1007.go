package main

import "fmt"

func F1007() {
	var a int
	var b int
	var c int
	var d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	diff := (a * b) - (c * d)
	fmt.Printf("DIFERENCA = %.d\n", diff)
}
