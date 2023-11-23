package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x float64
	y float64
}

func parseLine(line string) (float64, float64) {
	data := strings.Split(strings.TrimSpace(line), " ")
	print(data)
	x, _ := strconv.ParseFloat(data[0], 64)
	y, _ := strconv.ParseFloat(data[1], 64)

	return x, y
}

func F1015() {
	reader := bufio.NewReader(os.Stdin)
	var lines []string
	for i := 0; i < 2; i++ {
		line, _ := reader.ReadString('\n')
		lines = append(lines, line)
	}

	x1, y1 := parseLine(lines[0])
	x2, y2 := parseLine(lines[1])

	result := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	fmt.Printf("%.4f\n", result)

}
