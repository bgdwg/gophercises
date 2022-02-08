package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
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

func startQuizGame(problems [][]string) (int, error) {
	numCorrect := 0
	for i, problem := range problems {
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
	return numCorrect, nil
}

func main() {
	filenamePtr := flag.String("f", "problems.csv", "filename")
	flag.Parse()
	problems, err := readCsvFile(*filenamePtr)
	if err != nil {
		log.Fatal(err)
	}
	result, err := startQuizGame(problems)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("You scored %d out of %d.\n", result, len(problems))
}
