package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseScore(t *testing.T) {
	input := "Lions 3"

	parsed := parseScore(input)

	expected := score{
		Team:  "Lions",
		Goals: 3,
	}

	assert.Equal(t, expected, parsed)
}

func TestParseScoreWithMultipleSpaces(t *testing.T) {
	input := "Lions   2"

	parsed := parseScore(input)

	expected := score{
		Team:  "Lions",
		Goals: 2,
	}

	assert.Equal(t, expected, parsed)
}

func TestParseScoreWithSpaceInTeam(t *testing.T) {
	input := "Nelson Mandela Bay Lions 0"

	parsed := parseScore(input)

	expected := score{
		Team:  "Nelson Mandela Bay Lions",
		Goals: 0,
	}

	assert.Equal(t, expected, parsed)
}
