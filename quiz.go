package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type QA struct {
	Question string
	Answer   string
}

func main() {
	randomise := len(os.Args) > 1 && os.Args[1] == "r"
	quiz := loadQuiz(randomise)
	waitForUserToStart()

	totalCorrect, totalAsked := 0, 0
	go startTimer(30, len(quiz), &totalCorrect, &totalAsked)
	askQuestions(quiz, &totalCorrect, &totalAsked)
}

func loadQuiz(randomise bool) []QA {
	file, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(file))
	allRecords, _ := reader.ReadAll()

	quizQuestionsAndAnswers := make([]QA, len(allRecords))
	for i, record := range allRecords {
		quizQuestionsAndAnswers[i] = QA{record[0], record[1]}
	}

	if randomise {
		randomiseOrder(&quizQuestionsAndAnswers)
	}

	return quizQuestionsAndAnswers
}

func randomiseOrder(slice *[]QA) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*slice), func(i, j int) {
		(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
	})
}

func startTimer(duration time.Duration, totalPossible int, totalCorrect, totalAsked *int) {
	time.Sleep(duration * time.Second)
	fmt.Println("Times up!")
	printResults(totalPossible, totalCorrect, totalAsked)
}

func waitForUserToStart() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Press enter to start the quiz")
	reader.ReadString('\n')
}

func askQuestions(quiz []QA, totalCorrect, totalAsked *int) {
	// totalCorrect, totalAsked = 0, 0
	for _, qa := range quiz {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("\nWhat is %s\n", qa.Question)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		if text == qa.Answer {
			*totalCorrect++
		}
		*totalAsked++
	}
	printResults(len(quiz), totalCorrect, totalAsked)
}

func printResults(totalPossible int, totalCorrect, totalAsked *int) {
	fmt.Printf("\nYou managed to answer %d out of %d possible questions in 30 seconds.\n", *totalAsked, totalPossible)
	fmt.Printf("\nYour score is %d/%d.\n", *totalCorrect, *totalAsked)
	os.Exit(0)
}
