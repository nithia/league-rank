package league

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type RankTable map[string]*Team

func (r RankTable) Update(result Result) {
	s0 := result[0]
	s1 := result[1]

	t0 := r.getTeam(s0.Team)
	t1 := r.getTeam(s1.Team)

	if s0.Goals < s1.Goals {
		t0.Lost++
		t1.Won++
	} else if s0.Goals == s1.Goals {
		t0.Drawn++
		t1.Drawn++
	} else {
		t0.Won++
		t1.Lost++
	}

	t0.updatePoints()
	t1.updatePoints()
}

func (r RankTable) GetRankings() []*Team {
	ts := make([]*Team, 0, len(r))

	for _, t := range r {
		ts = append(ts, t)
	}

	sort.Slice(ts, func(i, j int) bool {
		if ts[i].Points != ts[j].Points {
			return ts[i].Points > ts[j].Points
		}
		return ts[i].Name < ts[j].Name
	})

	return ts

}

func (r RankTable) String() string {
	rankings := r.GetRankings()

	var b strings.Builder

	for i, team := range rankings {
		dups := 0

		if i > 0 && rankings[i-1].Points == team.Points {
			dups++
		} else {
			dups = 0
		}
		rank := i + 1 - dups
		_, _ = fmt.Fprintf(&b, "%d. %s, %d %s\n", rank, team.Name, team.Points, pts(team.Points))
	}
	return b.String()
}

func (r RankTable) getTeam(name string) *Team {
	if _, ok := r[name]; !ok {
		r[name] = &Team{
			Name:   name,
			Won:    0,
			Drawn:  0,
			Lost:   0,
			Points: 0,
		}
	}

	return r[name]
}

const (
	winPoints  = 3
	drawPoints = 1
	lossPoints = 0
)

func (t *Team) updatePoints() {
	t.Points = t.Won*winPoints + t.Drawn*drawPoints + t.Lost*lossPoints
}

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

func pts(n int) interface{} {
	if n == 1 {
		return "pt"
	}

	return "pts"
}
