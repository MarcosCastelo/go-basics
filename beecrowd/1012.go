package main

import "fmt"

func F1012() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	tri := (a * c) / 2
	cir := 3.14159 * (c * c)
	tra := ((a + b) * c) / 2
	qua := b * b
	ret := a * b

	fmt.Printf("TRIANGULO: %.3f\n", tri)
	fmt.Printf("CIRCULO: %.3f\n", cir)
	fmt.Printf("TRAPEZIO: %.3f\n", tra)
	fmt.Printf("QUADRADO: %.3f\n", qua)
	fmt.Printf("RETANGULO: %.3f\n", ret)
}
