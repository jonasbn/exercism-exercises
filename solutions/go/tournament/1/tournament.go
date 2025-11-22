package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type tournamentStats map[string]teamStats

type teamStats struct {
	name          string
	matchesPlayed uint8
	wins          uint8
	draws         uint8
	losses        uint8
	points        uint8
}

type ByPointsDesc []teamStats

func (a ByPointsDesc) Len() int { return len(a) }
func (a ByPointsDesc) Less(i, j int) bool {
	if a[i].points == a[j].points {
		return a[i].name < a[j].name
	} else {
		return a[i].points > a[j].points
	}
}
func (a ByPointsDesc) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Tally(reader io.Reader, writer io.Writer) error {

	tournament := tournamentStats{}

	hometeamStats := teamStats{}
	awayteamStats := teamStats{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		match := strings.Split(scanner.Text(), ";")

		if len(match) <= 1 {
			continue
		}

		if len(match) == 2 {
			return errors.New("unparsable data format, match not matching")
		}

		hometeam := match[0]
		awayteam := match[1]
		matchOutcome := match[2]

		if tournament[hometeam] != hometeamStats {
			hometeamStats = tournament[hometeam]
		}

		if tournament[awayteam] != awayteamStats {
			awayteamStats = tournament[awayteam]
		}

		switch matchOutcome {
		case "win":
			hometeamStats.wins += 1
			hometeamStats.points += 3
			awayteamStats.losses += 1
		case "draw":
			hometeamStats.draws += 1
			hometeamStats.points += 1
			awayteamStats.draws += 1
			awayteamStats.points += 1
		case "loss":
			hometeamStats.losses += 1
			awayteamStats.wins += 1
			awayteamStats.points += 3
		default:
			return errors.New("match outcome not recognized")
		}

		hometeamStats.matchesPlayed++
		awayteamStats.matchesPlayed++

		hometeamStats.name = hometeam
		awayteamStats.name = awayteam

		tournament[hometeam] = hometeamStats
		tournament[awayteam] = awayteamStats
	}

	teams := []teamStats{}

	for _, v := range tournament {
		teams = append(teams, v)
	}
	sort.Sort(ByPointsDesc(teams))

	fmt.Fprintf(writer, "%-[7]*[1]s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P", 30)
	for _, t := range teams {
		fmt.Fprintf(writer, "%-[7]*[1]s | %2d | %2d | %2d | %2d | %2d\n", t.name, t.matchesPlayed, t.wins, t.draws, t.losses, t.points, 30)
	}

	return nil
}
