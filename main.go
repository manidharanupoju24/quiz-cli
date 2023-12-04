package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello world")
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'") // csvFilename would be a pointer to the string
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename) // * to get the string and not the pointer
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided csv file")
	}
	// fmt.Println(lines)
	problems := parseLines(lines)
	//fmt.Println(problems)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			fmt.Println("Correct!")
			correct++
		} else {
			fmt.Println("Incorrect!")
		}
	}

	fmt.Printf("You got %d answers right and %d answers wrong!\n", correct, len(problems)-correct)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
