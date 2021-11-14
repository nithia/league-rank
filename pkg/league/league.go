package league

import (
	"strconv"
	"strings"
)

type RankTable []*Team

type Team struct {
	Name   string
	Won    int
	Drawn  int
	Lost   int
	Points int
}

type Score struct {
	Team  string
	Goals int
}

type Result []Score

func ParseResult(line string) Result {
	splitScores := strings.Split(line, ",")

	res := Result{}
	for _, splitScore := range splitScores {
		res = append(res, parseScore(strings.TrimSpace(splitScore)))
	}

	return res
}

func parseScore(input string) Score {
	lastSpace := strings.LastIndex(input, " ")

	return Score{
		Team:  strings.TrimSpace(input[:lastSpace]),
		Goals: parseInt(input[lastSpace+1:]),
	}
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s) // Ignore error as input is guaranteed to be well-formed

	return i
}
