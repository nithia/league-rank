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

func TestUpdateEmptyRankTable(t *testing.T) {
	table := RankTable{}

	input := "Lions 3, Tigers 1"
	result := ParseResult(input)

	table.Update(result)

	assert.Len(t, table, 2)
	assert.NotNil(t, table["Lions"])
	assert.NotNil(t, table["Tigers"])

	assert.Equal(t, 1, table["Lions"].Won)
	assert.Equal(t, 0, table["Lions"].Drawn)
	assert.Equal(t, 0, table["Lions"].Lost)
	assert.Equal(t, 3, table["Lions"].Points)

	assert.Equal(t, 0, table["Tigers"].Won)
	assert.Equal(t, 0, table["Tigers"].Drawn)
	assert.Equal(t, 1, table["Tigers"].Lost)
	assert.Equal(t, 0, table["Tigers"].Points)
}

func TestUpdateRankTable(t *testing.T) {
	table := RankTable{
		"Lions": {
			Name:   "Lions",
			Won:    1,
			Drawn:  0,
			Lost:   0,
			Points: 3,
		},
		"Tigers": {
			Name:   "Tigers",
			Won:    0,
			Drawn:  0,
			Lost:   1,
			Points: 0,
		},
	}

	input := "Tigers 1, Bears 2"
	result := ParseResult(input)

	table.Update(result)

	assert.Len(t, table, 3)
	assert.NotNil(t, table["Lions"])
	assert.NotNil(t, table["Tigers"])
	assert.NotNil(t, table["Bears"])

	assert.Equal(t, 1, table["Lions"].Won)
	assert.Equal(t, 0, table["Lions"].Drawn)
	assert.Equal(t, 0, table["Lions"].Lost)
	assert.Equal(t, 3, table["Lions"].Points)

	assert.Equal(t, 0, table["Tigers"].Won)
	assert.Equal(t, 0, table["Tigers"].Drawn)
	assert.Equal(t, 2, table["Tigers"].Lost)
	assert.Equal(t, 0, table["Tigers"].Points)

	assert.Equal(t, 1, table["Bears"].Won)
	assert.Equal(t, 0, table["Bears"].Drawn)
	assert.Equal(t, 0, table["Bears"].Lost)
	assert.Equal(t, 3, table["Bears"].Points)
}

func TestGetRankings(t *testing.T) {
	table := RankTable{
		"Lions": {
			Name:   "Lions",
			Won:    1,
			Drawn:  0,
			Lost:   0,
			Points: 3,
		},
		"Tigers": {
			Name:   "Tigers",
			Won:    0,
			Drawn:  1,
			Lost:   1,
			Points: 1,
		},
		"Bears": {
			Name:   "Bears",
			Won:    0,
			Drawn:  1,
			Lost:   0,
			Points: 1,
		},
	}

	rankings := table.GetRankings()

	assert.Len(t, rankings, 3)
	assert.Equal(t, "Lions", rankings[0].Name)
	assert.Equal(t, "Bears", rankings[1].Name)
	assert.Equal(t, "Tigers", rankings[2].Name)
}

func TestExample(t *testing.T) {
	input := []string{
		"Lions 3, Snakes 3",
		"Tarantulas 1, FC Awesome 0",
		"Lions 1, FC Awesome 1",
		"Tarantulas 3, Snakes 1",
		"Lions 4, Grouches 0"}

	expected := "1. Tarantulas, 6 pts\n2. Lions, 5 pts\n3. FC Awesome, 1 pt\n3. Snakes, 1 pt\n5. Grouches, 0 pts\n"

	table := RankTable{}
	for _, line := range input {
		table.Update(ParseResult(line))
	}

	rankings := table.String()

	assert.Equal(t, expected, rankings)
}

func TestUpdateAll(t *testing.T) {
	input := []string{
		"Lions 3, Snakes 3",
		"Tarantulas 1, FC Awesome 0",
		"Lions 1, FC Awesome 1",
		"Tarantulas 3, Snakes 1",
		"Lions 4, Grouches 0"}

	expected := "1. Tarantulas, 6 pts\n2. Lions, 5 pts\n3. FC Awesome, 1 pt\n3. Snakes, 1 pt\n5. Grouches, 0 pts\n"

	table := RankTable{}
	table.UpdateAll(input)

	rankings := table.String()

	assert.Equal(t, expected, rankings)
}
