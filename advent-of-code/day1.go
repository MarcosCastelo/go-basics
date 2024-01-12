package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

// On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

type StringEnd struct {
	position int
	char     string
}

var spelledNumbers = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func findStringNumberEnds(s string) (left StringEnd, right StringEnd) {
	left.position = len(s) - 1
	right.position = 0
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			left.char = string(s[i])
			left.position = i
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			right.char = string(s[i])
			right.position = i
			break
		}
	}

	return left, right
}

func findFisrtSpelledNumbers(s string) map[int]string {
	numbers := make(map[int]string)
	for index, word := range spelledNumbers {
		for i := 0; i < len(s)-len(word)+1; i++ {
			if s[i:i+len(word)] == word {
				numbers[i] = strconv.FormatInt(int64(index+1), 10)
			}
		}
	}
	return numbers
}

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

func main() {
	total := 0
	for _, line := range ReadLines("inputs/day1.txt") {
		l, r := findStringNumberEnds(line)
		leftSpelledNumbers := findFisrtSpelledNumbers(line[:l.position])
		if len(leftSpelledNumbers) != 0 {
			for k, v := range leftSpelledNumbers {
				if k < l.position {
					l.position = k
					l.char = v
				}
			}
		}
		rightSpelledNumbers := findFisrtSpelledNumbers(line[r.position:])
		if len(rightSpelledNumbers) != 0 {
			aux := r.position
			for k, v := range rightSpelledNumbers {
				if k+aux > r.position {
					r.position = k + aux
					r.char = v
				}
			}
		}
		num, _ := strconv.Atoi(l.char + r.char)
		fmt.Printf("%s  %d\n", line, num)
		fmt.Printf("left: %v, right: %v\n", l, r)
		total += num
	}
	fmt.Println("total: ", total)
}
