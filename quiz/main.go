package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	fmt.Println("Welcome to quiz")
	problemFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(problemFile))
	var quiz []Quiz
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
	totalQuestions := len(quiz)
	correct := 0
	ansReader := bufio.NewReader(os.Stdin)
	for i := 0; i < totalQuestions; i++ {
		fmt.Println(quiz[i].Question)
		ans, _ := ansReader.ReadString('\n')
		if strings.TrimRight(ans, "\n") == quiz[i].Answer {
			correct += 1
		}
	}
	fmt.Printf("%d correct out of %d \n", correct, totalQuestions)

}
