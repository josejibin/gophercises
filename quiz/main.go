package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

const quizTimeLimit = 6

func getQuiz(problemFileName string) []Quiz {
	var quiz []Quiz
	problemFile, _ := os.Open(problemFileName)
	reader := csv.NewReader(bufio.NewReader(problemFile))
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		quiz = append(quiz, Quiz{
			Question: line[0],
			Answer:   line[1],
		})
	}
	return quiz
}

func startQuiz(quiz []Quiz, totalQuestions int, answerQueue chan int) {
	correct := 0
	ansReader := bufio.NewReader(os.Stdin)
	for i := 0; i < totalQuestions; i++ {
		fmt.Println(quiz[i].Question)
		ans, _ := ansReader.ReadString('\n')
		if strings.TrimRight(ans, "\n") == quiz[i].Answer {
			correct += 1
		}
	}
	answerQueue <- correct
}

func main() {
	fmt.Println("Welcome to quiz")
	quiz := getQuiz("problems.csv")
	totalQuestions := len(quiz)
	answerQueue := make(chan int)
	timer1 := time.NewTimer(quizTimeLimit * time.Second)
	go startQuiz(quiz, totalQuestions, answerQueue)
loop:
	for {
		select {
		case correct := <-answerQueue:
			fmt.Printf("%d correct out of %d \n", correct, totalQuestions)
			break loop
		case <-timer1.C:
			break loop
		}
	}

}
