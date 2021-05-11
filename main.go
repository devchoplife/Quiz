package main

import (
	"flag"
	"fmt"
	"os"
)

//defining the type of datat in the csv file
type recordType struct {
	question string
	answer   string
}

func main() {
	//first we specify the filename
	fileName := flag.String("csv", "quiz.csv", "The csv file containing the questions and answers ")
	limit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	fkag.Parse()

	quizFile, err := os.Open(*fileName) //this opens the csv file

	//to display a message if  there is an error
	if err != nil {
		fmt.Printf("Failed to open the csv file: %s\n", *fileName)
		os.Exit(1)
	}
}
