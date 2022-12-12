package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type Buffer struct {
	data []byte
}

func (b *Buffer) Add(v byte) {
	b.data = append(b.data, v)
}

func (b Buffer) Check() bool {
	if len(b.data) < 14 {
		return false
	}
	set := map[byte]struct{}{}
	for _, c := range b.data[len(b.data)-14:] {
		set[c] = struct{}{}
	}
	return len(set) == 14
}

func main() {
	b := Buffer{}
	for i := range input {
		b.Add(input[i])
		if b.Check() {
			fmt.Println(i + 1)
			return
		}
	}
}
