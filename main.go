package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filePrt := flag.String("file", "problems.csv", "The Name of the Quiz File")
	timeLimit := flag.Int("time-limit", 30, "the time limit for the user to answer all problems")
	questions, err := os.Open(*filePrt)
	if err != nil {
		log.Fatal("Unable to parse file as CSV")
	}
	records, err := csv.NewReader(questions).ReadAll()

	//parse the lines in the inout file to structured
	probs := parseLines(records)
	fmt.Println("hey this is a test")
	runningTotal := 0
	//create a channel for the timer goroutine
	doneCh := make(chan bool)
	//loop over each problem in the parsed csv and prompt user for the answer
	go func() {
		for i, v := range probs {

			//display the problem for the user to solve
			fmt.Printf("Problem #%d: %s = \n", i+1, v.q)
			var answer string

			//Scanf takes user input and stores it in the pointer
			fmt.Scanf("%s\n", &answer)
			//if user input correct then increment running total
			if answer == v.a {
				runningTotal++
			}
		}
		doneCh <- true
	}()

	//select statement blocks code until there is data on either one of the channels.
	//so we are waiting fo doneCh or the timeC to produce output
	select {
	case <-doneCh:
		fmt.Println("You are done! good job!!!")

	case <-time.After(time.Duration(*timeLimit) * time.Second):
		fmt.Println("You have reached the maxium time")

	}

	fmt.Println("Your score:", runningTotal, "/", len(records))

	// go func() {
	// 	lol := time.NewTimer(time.Duration(time.Duration(*timeLimit) * time.Second))
	// 	<-lol.C
	// 	fmt.Println("")
	// 	fmt.Println("")
	// 	fmt.Println("Your time is up...")
	// }()

	fmt.Printf("You scored %d out of %d.\n", runningTotal, len(probs))
}

// var answer int = int(v.q)

// fmt.Println(probs)

func parseLines(records [][]string) []problem {
	//make a slice of type []problem with the size of len(records) called parsed problems
	parsedProblems := make([]problem, len(records))
	//loop over the csv object and put the values into a slice formatted in the struct type of problem
	for i, record := range records {
		parsedProblems[i] = problem{
			q: record[0],
			a: record[1],
		}

	}
	return parsedProblems
}

type problem struct {
	q string
	a string
}
