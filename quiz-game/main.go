package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Cannot open file %s", *csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Error parsing csv")
	}
	problems := ParseLines(lines)

	var correct int = 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d : %s = ?\n", i+1, problem.q)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == problem.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}

func ParseLines(lines [][]string) []problem {
	p := make([]problem, len(lines))
	for i, line := range lines {
		p[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return p
}

type problem struct {
	q string
	a string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
