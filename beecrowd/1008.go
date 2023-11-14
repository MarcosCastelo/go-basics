package main

import "fmt"

func F1008() {
	var id int
	var time int
	var salary float64

	fmt.Scan(&id)
	fmt.Scan(&time)
	fmt.Scan(&salary)

	fmt.Printf("NUMBER = %d\n", id)
	fmt.Printf("SALARY = U$ %.2f\n", float64(time)*salary)
}
