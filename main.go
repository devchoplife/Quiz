package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

//defining the type of datat in the csv file
type fileType struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

//THis function returns the struct fileType
func ParseRows(rows [][]string) []fileType {
	ini := make([]fileType, len(rows))
	for i, row := range rows {
		ini[i] = fileType{
			question: row[0],
			answer:   strings.TrimSpace(row[1]),
		}
	}
	return ini
}

func ReadStringInputWithTime(limit int) (string, error) {
	timer := time.NewTimer(time.Duration(limit) * time.Second).C
	done := make(chan bool)
	answer, err := "", error(nil)
	go func() {
		fmt.Scanf("%s\n", &answer)
		done <- true
	}()
	for {
		select {
		case <-timer:
			return "", errors.New("Timer expired")
		case <-done:
			return answer, err
		}
	}
}

func main() {
	//first we specify the filename
	fileName := flag.String("csv", "quiz.csv", "The csv file containing the questions and answers ")
	//then we specify the time limit for the quiz
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	quizFile, err := os.Open(*fileName) //this opens the csv file

	//to display a message if  there is an error
	if err != nil {
		fmt.Printf("Failed to open the csv file: %s\n", *fileName)
		os.Exit(1)
	}

	defer quizFile.Close() //close the csv file

	readerQuiz := csv.NewReader(quizFile) //function to read the csv file
	rows, err := readerQuiz.ReadAll()     //specify the number of rows to be read in
	if err != nil {
		//This errpr is thrown if the CSV parse was not successful
		exit("Failed to parse the provided CSV file")
	}

	//Parse the rows using a function
	quiz := ParseRows(rows)

	correctAnsCount := 0
	for i, p := range quiz {
		fmt.Printf("Quiz No. %d: %s =", i+1, p.question)

		answer, err := ReadStringInputWithTime(*limit)
		if err != nil {
			println("Time Expired!")
			break
		}
		if strings.ToLower(strings.Trim(answer, "\n ")) == p.answer {
			correctAnsCount++
		}
	}
	println("You scored ", correctAnsCount, "out of ", len(quiz))
}
