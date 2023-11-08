package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"
)

type Question struct {
	question string
	answer   string
	marked   bool
}

func main() {
	var filename string
	var shuffle bool
	var countdown int

	flag.StringVar(&filename, "filename", "problems.csv", "pass the csv file to load problems (Default: problems.csv)")
	flag.BoolVar(&shuffle, "shuffle", false, "set true or false to shuffle the problems (Default: false)")
	flag.IntVar(&countdown, "time", 30, "set the time for answer all questions (Default: 30)")

	flag.Parse()

	score := 0
	questions, questions_ids := getQuestionsFromCSV(filename)

	startTime := time.Now()

	index := 0
	id := questions_ids[index]
	remainingTime := 0
	answerCount := 0

	fmt.Println("Welcome to quizz game now answer a series of questions and get the higher score")
	for range time.Tick(1 * time.Second) {
		remainingTime = countdown - decreaseTime(startTime)
		if remainingTime <= 0 {
			break
		}
		if len(questions_ids) < 1 {
			break
		}
		fmt.Print("\033[H\033[2J")
		fmt.Println("Time reaming: \t\t\t", remainingTime)
		if shuffle == true {
			questionsLength := len(questions_ids)
			index = rand.Intn(questionsLength)
			id = questions_ids[index]
			questions_ids = append(questions_ids[:index], questions_ids[index+1:]...)
		}
		score += scoreQuestion(&questions[id])
		answerCount = len(questions_ids)
		index += 1
	}

	if answerCount > 0 {
		fmt.Println("Close!, reamaning questions: ", answerCount)
	} else {
		fmt.Println("Congratulations! you finished")
	}
	fmt.Println("Your final score is: ", score)
}

func getQuestionsFromCSV(filename string) ([]Question, []int) {
	var questions []Question
	var questions_ids []int
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error readings records!")
	}

	for index, each_record := range records {
		questions_ids = append(questions_ids, index)
		questions = append(questions, Question{
			question: each_record[0],
			answer:   each_record[1],
			marked:   false,
		})
	}
	return questions, questions_ids
}

func scoreQuestion(question *Question) int {
	fmt.Printf("%s: ", question.question)
	var answer string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	answer = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(scanner.Text(), "")
	if answer == question.answer {
		return 1
	}
	return 0
}

func decreaseTime(t time.Time) int {
	currentTime := time.Now()
	difference := t.Sub(currentTime)
	return -int(difference.Seconds())
}
