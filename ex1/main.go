package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

func main() {
	csvFileName := flag.String("file", "ex1/test.csv", "a file in csv format")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	flag.Parse()

	if *csvFileName == "" {
		log.Fatal("Please provide a CSV file using -file flag")
	}

	lines, err := ReadCSVFile(*csvFileName)
	if err != nil {
		log.Fatalf("Failed to parse the provided CSV file: %s", err)
	}

	problems := ParseLines(lines)
	if len(problems) == 0 {
		log.Fatal("No problems found in the CSV file")
	}

	correct := 0
	answerCh := make(chan string)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	go readInput(answerCh)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s \n", i+1, p.Question)
		select {
		case <-timer.C:
			fmt.Println("Time's up!")
			fmt.Printf("You scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if strings.TrimSpace(answer) == p.Answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func ParseLines(lines [][]string) []Problem {
	r := make([]Problem, len(lines))
	for i, line := range lines {
		r[i] = Problem{
			Question: strings.TrimSpace(line[0]),
			Answer:   strings.TrimSpace(line[1]),
		}
	}
	return r
}

func ReadCSVFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func readInput(answerCh chan string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answerCh <- scanner.Text()
	}
	if scanner.Err() != nil {
		log.Fatalf("Failed to read input: %s", scanner.Err())
	}
	close(answerCh)
}
