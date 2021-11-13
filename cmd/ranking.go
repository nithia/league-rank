package main

import (
	"strconv"
	"strings"
)

type score struct {
	Team  string
	Goals int
}

type result []score

func main() {

}

func parseResult(line string) result {
	splitScores := strings.Split(line, ",")

	res := result{}
	for _, splitScore := range splitScores {
		res = append(res, parseScore(strings.TrimSpace(splitScore)))
	}

	return res
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
