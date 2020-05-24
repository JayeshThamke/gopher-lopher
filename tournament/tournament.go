package tournament

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"
)

// MatchResult a type to represent end result of match
type MatchResult int

func getMatchResult(result string) MatchResult {
	result = strings.ToLower(result)
	switch result {
	case "win":
		return MatchResult(3)
	case "loss":
		return MatchResult(0)
	case "draw":
		return MatchResult(1)
	default:
		return MatchResult(-1)
	}
}

// Keeping the state of statistics of each team here
// like in memory db or k-v pair
var teamsMap = make(map[string]*Team)

// Team respresnts a team in a football match
type Team struct {
	Name  string
	Stats *Statistics
}

// NewTeam is a constructor function for Team
func NewTeam(name string) *Team {
	return &Team{
		Name:  name,
		Stats: &Statistics{},
	}
}

// Statistics represents the score and
type Statistics struct {
	MatchesPlayed int
	MatchesWon    int
	MatchesLost   int
	MatchesDrawn  int
	TotalPoints   int
}

// Match represents the match (event) between two teams
type Match struct {
	TeamA, TeamB Team
	Result       MatchResult
}

// Tally tallys the tournament summary
func Tally(r io.Reader, w io.Writer) error {
	in := streamToString(r)

	matches := strings.Split(in, "\n")
	for _, match := range matches {
		m := strings.Split(match, ";")

		footballMatch, err := validateMatch(m)
		if err != nil {
			return err
		}

		err = updateStats(&footballMatch)
		if err != nil {
			return err
		}
	}

	var teams []*Team = []*Team{}

	for _, t := range teamsMap {
		calculateTotalPoints(t)
		teams = append(teams, t)
	}

	// sorting in descending points list as per requirement
	sort.SliceStable(teams, func(i, j int) bool {
		return teams[i].Stats.TotalPoints > teams[j].Stats.TotalPoints
	})

	var b bytes.Buffer
	err := formatStatistics(&b, teams)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(w, &b)
	if err != nil {
		return err
	}
	return nil
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	input := buf.String()

	// strip leading and trailing new line that
	// was made easy for reading, it helps to split
	// tournament string into matches and further details.
	input = input[1 : len(input)-1]
	return input
}

func validateMatch(m []string) (Match, error) {
	if len(m) != 3 {
		return Match{}, fmt.Errorf("Insufficent match details. Required elements 3 but found %d", len(m))
	}

	// loop is only for teams
	for _, teamName := range m[:len(m)-1] {
		if _, teamExists := teamsMap[teamName]; !teamExists {
			teamsMap[teamName] = NewTeam(teamName)
		}
	}

	match := Match{
		TeamA:  *teamsMap[m[0]],
		TeamB:  *teamsMap[m[1]],
		Result: getMatchResult(m[2]),
	}

	if match.Result == MatchResult(-1) {
		return Match{}, fmt.Errorf("Invalid match result found")
	}

	return match, nil
}

func updateStats(match *Match) error {
	// update MatchesPlayed
	match.TeamA.Stats.MatchesPlayed++
	match.TeamB.Stats.MatchesPlayed++

	switch match.Result {
	// when match is won by teamA
	case MatchResult(3):
		match.TeamA.Stats.MatchesWon++
		match.TeamB.Stats.MatchesLost++
	// when match is lost by teamA
	case MatchResult(0):
		match.TeamA.Stats.MatchesLost++
		match.TeamB.Stats.MatchesWon++
	// when match is drawn
	case MatchResult(1):
		match.TeamA.Stats.MatchesDrawn++
		match.TeamB.Stats.MatchesDrawn++
	}

	return nil
}

func calculateTotalPoints(team *Team) error {
	const (
		winPt  = 3
		lostPt = 0
		drawPt = 1
	)

	team.Stats.TotalPoints = team.Stats.MatchesWon*winPt +
		team.Stats.MatchesLost*lostPt +
		team.Stats.MatchesDrawn*drawPt

	return nil
}

func formatStatistics(b *bytes.Buffer, teams []*Team) error {
	_, err := b.WriteString(fmt.Sprintf("%-31v| MP |  W |  D |  L |  P", "Team"))
	if err != nil {
		return err
	}

	for _, team := range teams {
		_, err := b.WriteString(fmt.Sprintf("\n%-31v|  %d |  %d |  %d |  %d |  %d", team.Name,
			team.Stats.MatchesPlayed,
			team.Stats.MatchesWon,
			team.Stats.MatchesDrawn,
			team.Stats.MatchesLost,
			team.Stats.TotalPoints))

		if err != nil {
			return err
		}
	}
	return nil
}
