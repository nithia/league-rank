package main

import (
	"strconv"
	"strings"
)

type score struct {
	Team  string
	Goals int
}

type result [2]score

func main() {

}

func parseResult(line string) result {
	return [2]score{}
}

func parseScore(input string) score {
	lastSpace := strings.LastIndex(input, " ")

	return score{
		Team:  strings.TrimSpace(input[:lastSpace]),
		Goals: parseInt(input[lastSpace+1:]),
	}
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s) // Ignore error as input is guaranteed to be well-formed

	return i
}
