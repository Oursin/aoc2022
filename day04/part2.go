package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Assignment struct {
	start int
	end   int
}

func (a Assignment) TotalOverlap(b Assignment) bool {
	return a.start <= b.start && a.end >= b.end ||
		a.start >= b.start && a.end <= b.end
}

func (a Assignment) PartialOverlap(b Assignment) bool {
	return a.start >= b.start && a.start <= b.end || a.end >= b.start && a.end <= b.end
}

func parseAssignment(input string) Assignment {
	var assignment Assignment
	values := strings.Split(input, "-")
	assignment.start, _ = strconv.Atoi(values[0])
	assignment.end, _ = strconv.Atoi(values[1])
	return assignment
}

func main() {
	count := 0
	lines := strings.Split(input, "\n")
	for _, line := range lines[:len(lines)-1] {
		pair := strings.Split(line, ",")
		assignment1 := parseAssignment(pair[0])
		assignment2 := parseAssignment(pair[1])
		fmt.Println(assignment1, assignment2)
		if assignment1.PartialOverlap(assignment2) || assignment1.TotalOverlap(assignment2) {
			fmt.Println("true")
			count++
		}
	}
	fmt.Println(count)
}
