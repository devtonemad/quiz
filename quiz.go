package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Welcome to the quiz \n")
	csvFilename := flag.String("csv", "problems.csv", "a csv file int the format 'question,answer'")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	_ = file
	if err != nil {
		exit(fmt.Sprintf("Problem when reading file name %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	_ = r
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse provided csv file")
	}

	fmt.Printf("Lines: %s\n", lines)

	problems := parseLines(lines)
	fmt.Printf("Lines as array of problems %s\n", parseLines(lines))

	correct := 0
	for i, p := range problems {
		fmt.Printf("%d. %s:\n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}

	fmt.Printf("You result is %d out of %d\n", correct, len(problems))

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

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
