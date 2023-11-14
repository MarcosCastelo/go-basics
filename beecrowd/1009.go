package main

import "fmt"

func F1009() {
	var name string
	var salary float64
	var sells float64

	fmt.Scan(&name)
	fmt.Scan(&salary)
	fmt.Scan(&sells)

	fmt.Printf("TOTAL = R$ %.2f\n", salary+(0.15*sells))
}
