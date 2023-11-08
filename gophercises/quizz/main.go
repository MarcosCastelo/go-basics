package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
)

type Question struct {
	question string
	answer   string
	marked   bool
}

func main() {
	score := 0
	questions, questions_ids := getQuestionsFromCSV("problems.csv")
	fmt.Println("Welcome to quizz game now answer a series of questions and get the higher score")
	for range questions_ids {
		index := rand.Intn(len(questions_ids))
		id := questions_ids[index]
		score += scoreQuestion(&questions[id])
		questions_ids = append(questions_ids[:index], questions_ids[index+1:]...)
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
	var answer string
	fmt.Printf("%s: ", question.question)
	_, err := fmt.Scanln(&answer)
	if err != nil {
		log.Fatal(err)
	}
	if answer == question.answer {
		return 1
	}
	return 0
}
