package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	tasks := readcsv("problems.csv")
	answers := make(map[string]int)
	for task := range tasks {
		answers[task] = requestAnswer(task + ": ")
	}
	score, total := scoreAndTotal(tasks, answers)
	fmt.Println("Your score is: ", score, "/", total)
}

func requestAnswer(message string) int {
	var answer int
	fmt.Print(message)
	_, err := fmt.Scanf("%d", &answer)
	if err != nil {
		log.Fatal("Error reading")
	}
	return answer
}

func scoreAndTotal(correct, attempted map[string]int) (int, int) {
	score := 0
	total := len(correct)
	for task := range correct {
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
