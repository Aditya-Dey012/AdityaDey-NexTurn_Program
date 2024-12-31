package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	QuestionText string
	Options      [4]string
	CorrectAns   int
}

var questionBank = []Question{
	{
		QuestionText: "What is the capital of india?",
		Options:      [4]string{"1. Assam", "2. Delhi", "3. Hyderabad", "4. Gujarat"},
		CorrectAns:   2,
	},
	{
		QuestionText: "Who is the current Prime Minister of India",
		Options:      [4]string{"1. Narendra Modi", "2. Rahul Gandhi", "3. Jawaharlal Nehru", "4. MS Dhoni"},
		CorrectAns:   1,
	},
	{
		QuestionText: "Who is a Cricketer?",
		Options:      [4]string{"1. Pele", "2. Virat Kholi", "3. Ronaldo", "4. Neymar"},
		CorrectAns:   2,
	},
}

const questionTimeLimit = 10 * time.Second

func TakeQuiz() {
	var score int
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questionBank {
		fmt.Printf("Question %d: %s\n", i+1, question.QuestionText)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		answerChan := make(chan int)
		timer := time.NewTimer(questionTimeLimit)

		go func() {
			for {
				fmt.Print("Enter your answer (or type 'exit' to quit): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if input == "exit" {
					close(answerChan)
					return
				}

				answer, err := strconv.Atoi(input)
				if err != nil || answer < 1 || answer > 4 {
					fmt.Println("Invalid input. Please enter the correct option.")
					continue
				}
				answerChan <- answer
				return
			}
		}()

		select {
		case answer, ok := <-answerChan:
			if !ok {
				fmt.Println("Exiting quiz........")
				return
			}
			if answer == question.CorrectAns {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Incorrect")
			}
		case <-timer.C:
			fmt.Println("Tick-Tock ; Time's up")
		}
		fmt.Println()
	}

	fmt.Printf("Quiz completed! Your score is: %d/%d\n", score, len(questionBank))
	if score == len(questionBank) {
		fmt.Println("Performance: Excellent")
	} else if score >= len(questionBank)/2 {
		fmt.Println("Performance: Good")
	} else {
		fmt.Println("Performance: Do Better")
	}
}

func main() {
	fmt.Println("Online Examination System!")

	TakeQuiz()
}
