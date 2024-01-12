package main

import (
	"fmt"
)

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int64
}

func main() {
	p1 := new(Pessoa)
	p2 := new(Pessoa)

	p1.Nome = "Marcos"
	p2.Nome = "Vitor"

	fmt.Println(p1, p2)
}
