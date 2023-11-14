package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Sale struct {
	id     int
	nItems int
	price  float64
}

func setSale(line string) Sale {
	data := strings.Split(strings.TrimSpace(line), " ")
	id, _ := strconv.Atoi(data[0])
	nItems, _ := strconv.Atoi(data[1])
	price, _ := strconv.ParseFloat(data[2], 64)

	return Sale{
		id:     id,
		nItems: nItems,
		price:  price,
	}
}

func F1010() {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for i := 0; i < 2; i++ {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}

	sale1 := setSale(lines[0])
	sale2 := setSale(lines[1])

	total := (float64(sale1.nItems) * sale1.price) + (float64(sale2.nItems) * sale2.price)
	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
