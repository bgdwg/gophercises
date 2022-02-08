package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"reflect"
)

func main() {
	filenamePtr := flag.String("f", "problems.csv", "filename")
	flag.Parse()
	f, err := os.Open(*filenamePtr)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	cntCorrect := 0
	for i, record := range records {
		fmt.Printf("Question %d: %s=", i+1, record[0])
		reader := bufio.NewReader(os.Stdin)
		correctAns := []byte(record[1])
		ans := make([]byte, len(correctAns))
		_, err := reader.Read(ans)
		if err != nil {
			panic(err)
		}
		if reflect.DeepEqual(ans, correctAns) {
			cntCorrect++
		} else {
			break
		}
	}
	fmt.Printf("Result: %d out of %d\n", cntCorrect, len(records))
}
