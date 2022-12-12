package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) PushN(v []T) {
	s.data = append(s.data, v...)
}

func (s *Stack[T]) Pop() T {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack[T]) PopN(n int) []T {
	v := s.data[len(s.data)-n : len(s.data)]
	s.data = s.data[:len(s.data)-n]
	return v

}

type Stacks []Stack[string]

func parseStacks(lines []string) Stacks {
	stackList := lines[len(lines)-1]
	totalStacks, _ := strconv.Atoi(string(stackList[len(stackList)-2]))
	stacks := make(Stacks, totalStacks)
	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		stackindex := 0
		for j := 1; j <= len(line); j += 4 {
			if unicode.IsLetter(rune(line[j])) {
				stacks[stackindex].Push(string(line[j]))
			}
			stackindex++
		}
	}
	return stacks
}

type Move struct {
	nb, src, dest int
}

func (m Move) Execute(s Stacks) {
	if (m.src) == 0 || m.dest == 0 {
		return
	}
	src := &s[m.src-1]
	dest := &s[m.dest-1]
	dest.PushN(src.PopN(m.nb))
}

func parseMoves(lines []string) []Move {
	moves := make([]Move, len(lines))
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		var move Move
		fmt.Sscanf(line, "move %d from %d to %d", &move.nb, &move.src, &move.dest)
		moves[i] = move
	}

	return moves
}

func BeTheCrane(stacks Stacks, moves []Move) {
	fmt.Println(stacks)
	for _, move := range moves {
		fmt.Println(move)
		move.Execute(stacks)
		fmt.Println(stacks)
	}
}

func main() {
	parts := strings.Split(input, "\n\n")
	stacks := parseStacks(strings.Split(parts[0], "\n"))
	moves := parseMoves(strings.Split(parts[1], "\n"))
	BeTheCrane(stacks, moves)
	fmt.Println(stacks)
	for _, stack := range stacks {
		fmt.Printf("%s", stack.Pop())
	}
	fmt.Printf("\n")
}
