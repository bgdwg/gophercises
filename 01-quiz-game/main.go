package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

func readCsvFile(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, fmt.Errorf("unable to open %s - %w", filename, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, fmt.Errorf("unable to read all records from %s - %w", filename, err)
	}
	return records, nil
}

func startQuizGame(problems [][]string, timeLimit time.Duration) (int, error) {
	fmt.Println("Are you ready? Press enter to start quiz...")
	reader := bufio.NewReader(os.Stdin)
	if confirm, err := reader.ReadByte(); err != nil || confirm != '\n' {
		return 0, fmt.Errorf("invalid key pressed")
	}
	fmt.Println(timeLimit)
	timer := time.NewTimer(timeLimit * time.Second)

	numCorrect := 0
	for i, problem := range problems {
		select {
		case <-timer.C:
			return numCorrect, nil
		default:
			question, correctAnswer := problem[0], []byte(problem[1])
			fmt.Printf("Problem #%d: %s = ", i+1, question)
			reader := bufio.NewReader(os.Stdin)
			answer := make([]byte, len(correctAnswer))
			if _, err := reader.Read(answer); err != nil {
				return 0, fmt.Errorf("unable to read your answer")
			}
			if !reflect.DeepEqual(answer, correctAnswer) {
				break
			}
			numCorrect++
		}
	}
	return numCorrect, nil
}

func main() {
	filenamePtr := flag.String("file", "problems.csv", "filename")
	timeLimitPtr := flag.Int("limit", 30, "time limit")
	flag.Parse()

	problems, err := readCsvFile(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}
	result, err := startQuizGame(problems, time.Duration(*timeLimitPtr))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You scored %d out of %d.\n", result, len(problems))
}
