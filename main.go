package main

import (
	"encoding/csv"
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
	os.Exit(1) //return error
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

func main() {
	//first we specify the filename
	fileName := flag.String("csv", "quiz.csv", "The csv file containing the questions and answers ")
	//then we specify the time limit for the quiz
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse() /*Parse parses the command-line flags from os.Args[1:].
	Must be called after all flags are defined and before
	flags are accessed by the program.*/

	quizFile, err := os.Open(*fileName) //this opens the csv file

	//to display a message if  there is an error
	if err != nil {
		fmt.Printf("Failed to open the csv file: %s\n", *fileName)
		os.Exit(1) /*Exit causes the current program to exit with the given
		status code. Conventionally, code zero indicates success, non-zero an
		error. The program terminates immediately; deferred functions
		are not run.*/
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

quizLoop:
	for i, p := range quiz {
		fmt.Printf("Quiz No. %d: %s =", i+1, p.question)

		timer := time.NewTimer(time.Duration(*limit) * time.Second) //sends the timer to the Channel
		var answer string
		answerCh := make(chan string)

		//go routine
		go func() {
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer //this is the answer channel
		}()

		select {
		case <-timer.C:
			fmt.Println("Time Expired!!!")
			break quizLoop //breaks the Loop

		case answer := <-answerCh:
			if answer == p.answer {
				correctAnsCount++
			}
		}
	}
	fmt.Println("You scored ", correctAnsCount, "out of ", len(quiz))
}
