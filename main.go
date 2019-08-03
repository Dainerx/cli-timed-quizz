package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const QUIZ_DURATION = 60

func main() {
	fmt.Println("Welcome to CLI Timed Quiz!")
	fmt.Println("You have ", QUIZ_DURATION, " seconds to finish the quiz")
	timer := time.NewTimer(time.Second * QUIZ_DURATION)
	tasks := readcsv("problems.csv")
	answers := make(map[string]int)
	for task := range tasks {
		answerChannel := make(chan int)
		go func() {
			answers[task] = requestAnswer(task + ": ")
			answerChannel <- answers[task]
		}()
		select {
		case <-timer.C:
			score, total := scoreAndTotal(tasks, answers)
			fmt.Println("Your score is: ", score, "/", total)
			return
		case <-answerChannel:
			continue
		}
	}
	score, total := scoreAndTotal(tasks, answers)
	fmt.Println("Your score is: ", score, "/", total)
}

func requestAnswer(message string) int {
	var answer int
	fmt.Print(message)
	_, err := fmt.Scanf("%d", &answer) //blocking call
	if err != nil {
		log.Fatal("Error reading")
	}
	return answer
}

func scoreAndTotal(correct, attempted map[string]int) (int, int) {
	score := 0
	total := len(correct)
	for task := range correct {
		if _, ok := attempted[task]; !ok {
			continue
		}
		if correct[task] == attempted[task] {
			score++
		}
	}
	return score, total
}

func readcsv(fileName string) map[string]int {
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error opening the file")
	}

	tasks := make(map[string]int)
	r := csv.NewReader(fileReader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error reading record from file")
		}

		tasks[record[0]], _ = strconv.Atoi((record[1]))
	}

	return tasks
}
