package league

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseScore(t *testing.T) {
	input := "Lions 3"

	parsed := parseScore(input)

	expected := Score{
		Team:  "Lions",
		Goals: 3,
	}

	assert.Equal(t, expected, parsed)
}

func TestParseScoreWithMultipleSpaces(t *testing.T) {
	input := "Lions   2"

	parsed := parseScore(input)

	expected := Score{
		Team:  "Lions",
		Goals: 2,
	}

	assert.Equal(t, expected, parsed)
}

func TestParseScoreWithSpaceInTeam(t *testing.T) {
	input := "Nelson Mandela Bay Lions 0"

	parsed := parseScore(input)

	expected := Score{
		Team:  "Nelson Mandela Bay Lions",
		Goals: 0,
	}

	assert.Equal(t, expected, parsed)
}

func TestParseResult(t *testing.T) {
	input := "Lions 3, Tigers 1"

	parsed := ParseResult(input)

	expected := Result{
		Score{
			Team:  "Lions",
			Goals: 3,
		},
		Score{
			Team:  "Tigers",
			Goals: 1,
		},
	}

	assert.Equal(t, expected, parsed)
}

func TestParseResultWithExtraSpaces(t *testing.T) {
	input := " Nelson Mandela Bay Lions 1 ,   Cape Town Tigers 2 "

	parsed := ParseResult(input)

	expected := Result{
		Score{
			Team:  "Nelson Mandela Bay Lions",
			Goals: 1,
		},
		Score{
			Team:  "Cape Town Tigers",
			Goals: 2,
		},
	}

	assert.Equal(t, expected, parsed)
}
