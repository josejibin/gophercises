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

func startQuiz(quiz []Quiz, totalQuestions int, questionQueue chan Quiz) {
	for i := 0; i < totalQuestions; i++ {
		questionQueue <- quiz[i]
	}
}

func main() {
	fmt.Println("Welcome to quiz")
	quiz := getQuiz("problems.csv")
	correct := 0
	attempted := 0
	ansReader := bufio.NewReader(os.Stdin)
	totalQuestions := len(quiz)
	questionQueue := make(chan Quiz)
	timer1 := time.NewTimer(quizTimeLimit * time.Second)
	go startQuiz(quiz, totalQuestions, questionQueue)
loop:
	for {
		select {
		case q := <-questionQueue:
			fmt.Println(q.Question)
			ans, _ := ansReader.ReadString('\n')
			if strings.TrimRight(ans, "\n") == q.Answer {
				correct += 1
			}
			attempted += 1
			if attempted == totalQuestions {
				fmt.Printf("%d correct out of %d \n", correct, totalQuestions)
				break loop
			}
		case <-timer1.C:
			fmt.Println("Time ended")
			fmt.Printf("%d correct out of %d \n", correct, totalQuestions)
			break loop
		}
	}

}
