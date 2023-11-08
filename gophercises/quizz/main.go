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
)

type Question struct {
	question string
	answer   string
	marked   bool
}

func main() {
	var filename string
	var shuffle bool
	var time int

	flag.StringVar(&filename, "filename", "problems.csv", "pass the csv file to load problems (Default: problems.csv)")
	flag.BoolVar(&shuffle, "shuffle", false, "set true or false to shuffle the problems (Default: false)")
	flag.IntVar(&time, "time", 30, "set the time for answer all questions (Default: 30)")

	flag.Parse()

	score := 0
	questions, questions_ids := getQuestionsFromCSV(filename)
	fmt.Println("Welcome to quizz game now answer a series of questions and get the higher score")
	for index, id := range questions_ids {
		if shuffle == true {
			index = rand.Intn(len(questions_ids))
			id = questions_ids[index]
			questions_ids = append(questions_ids[:index], questions_ids[index+1:]...)
		}
		score += scoreQuestion(&questions[id])
	}
	fmt.Println("Congratulations, your final score is ", score)
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
	reader := bufio.NewReader(os.Stdin)
	answer, err := reader.ReadString('\n')
	answer = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(answer, "")
	if err != nil {
		log.Fatal(err)
	}
	if answer == question.answer {
		return 1
	}
	return 0
}
