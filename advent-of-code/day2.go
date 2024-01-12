package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) (lines []string) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text() // GET the line string
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	return lines
}

func fetchLine(line string) (int, bool) {
	s := strings.Split(line, ":")

	gameId, _ := strconv.Atoi(s[0][5:])

	// rounds := strings.FieldsFunc()(s[1], ";")

	return gameId, true
}

func checkBag(bag map[string]int) bool {
	return bag["red"] <= 12 && bag["green"] <= 13 && bag["blue"] <= 14
}

func main() {
	gameIdsSum := 0
	for _, line := range ReadLines("inputs/day2.txt") {
		gameId, ok := fetchLine(line)
		if ok {
			gameIdsSum += gameId
		}
	}
	fmt.Println(len(ReadLines("inputs/day2.txt")))
	// fmt.Println("Total: ", gameIdsSum)

}
