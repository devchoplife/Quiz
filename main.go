package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	quizFile := flag.String("csv", "quiz.csv", "csv file containing the quiz questions and answers")
	flag.Parse()

	file, err := os.Open(*quizFile)
	if err != nil {
		fmt.Printf("Failed to open the csv file: %s\n", *csvFile)
		os.Exit(1)
		//the above can also be written as
		//exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *quizFile))
	}

	defer quizFile.Close() //close the csv file

	quiz := csv.NewReader(quizFile)
	rows, err := quiz.ReadAll()

	if err != nil {
		exit("Failed to parse the provided dsv file")
	}

	issues := ParseLines(rows)

	correctAnswerNum := 0
	for i, p := range issues {
		fmt.Printf("Problem #%d %s = ", i+1, p.question)

		ans, err := ReadStringWithLithLimitTime(^limit)
	}
}
