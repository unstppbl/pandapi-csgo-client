package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Match struct {
	BeginAt       time.Time   `json:"begin_at"`
	DetailedStats bool        `json:"detailed_stats"`
	Draw          bool        `json:"draw"`
	EndAt         interface{} `json:"end_at"`
	Forfeit       bool        `json:"forfeit"`
	GameAdvantage interface{} `json:"game_advantage"`
	Games         []struct {
		BeginAt       interface{} `json:"begin_at"`
		DetailedStats bool        `json:"detailed_stats"`
		EndAt         interface{} `json:"end_at"`
		Finished      bool        `json:"finished"`
		Forfeit       bool        `json:"forfeit"`
		ID            int         `json:"id"`
		Length        interface{} `json:"length"`
		MatchID       int         `json:"match_id"`
		Position      int         `json:"position"`
		Status        string      `json:"status"`
		VideoURL      interface{} `json:"video_url"`
		Winner        struct {
			ID   interface{} `json:"id"`
			Type string      `json:"type"`
		} `json:"winner"`
		WinnerType string `json:"winner_type"`
	} `json:"games"`
	ID     int `json:"id"`
	League struct {
		ID         int         `json:"id"`
		ImageURL   string      `json:"image_url"`
		ModifiedAt time.Time   `json:"modified_at"`
		Name       string      `json:"name"`
		Slug       string      `json:"slug"`
		URL        interface{} `json:"url"`
	} `json:"league"`
	LeagueID int `json:"league_id"`
	Live     struct {
		OpensAt   interface{} `json:"opens_at"`
		Supported bool        `json:"supported"`
		URL       interface{} `json:"url"`
	} `json:"live"`
	StreamsList []struct {
		EmbedURL string `json:"embed_url"`
		RawURL   string `json:"raw_url"`
		Language string `json:"language"`
		Official bool   `json:"official"`
		Main     bool   `json:"main"`
	} `json:"streams_list"`
	MatchType     string    `json:"match_type"`
	ModifiedAt    time.Time `json:"modified_at"`
	Name          string    `json:"name"`
	NumberOfGames int       `json:"number_of_games"`
	Opponents     []struct {
		Opponent struct {
			ID       int    `json:"id"`
			Location string `json:"location"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
		} `json:"opponent"`
		Type string `json:"type"`
	} `json:"opponents"`
	Rescheduled bool `json:"rescheduled"`
	Results     []struct {
		Score  int `json:"score"`
		TeamID int `json:"team_id"`
	} `json:"results"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Serie       struct {
		BeginAt     time.Time   `json:"begin_at"`
		Description interface{} `json:"description"`
		EndAt       interface{} `json:"end_at"`
		FullName    string      `json:"full_name"`
		ID          int         `json:"id"`
		LeagueID    int         `json:"league_id"`
		ModifiedAt  time.Time   `json:"modified_at"`
		Name        string      `json:"name"`
		Season      interface{} `json:"season"`
		Slug        string      `json:"slug"`
		WinnerID    interface{} `json:"winner_id"`
		WinnerType  interface{} `json:"winner_type"`
		Year        int         `json:"year"`
	} `json:"serie"`
	SerieID    int    `json:"serie_id"`
	Slug       string `json:"slug"`
	Status     string `json:"status"`
	Tournament struct {
		BeginAt       time.Time   `json:"begin_at"`
		EndAt         interface{} `json:"end_at"`
		ID            int         `json:"id"`
		LeagueID      int         `json:"league_id"`
		LiveSupported bool        `json:"live_supported"`
		ModifiedAt    time.Time   `json:"modified_at"`
		Name          string      `json:"name"`
		Prizepool     interface{} `json:"prizepool"`
		SerieID       int         `json:"serie_id"`
		Slug          string      `json:"slug"`
		WinnerID      interface{} `json:"winner_id"`
		WinnerType    interface{} `json:"winner_type"`
	} `json:"tournament"`
	TournamentID int `json:"tournament_id"`
	Videogame    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"videogame"`
	VideogameVersion interface{} `json:"videogame_version"`
	Winner           interface{} `json:"winner"`
	WinnerID         interface{} `json:"winner_id"`
}

func ParseMatches(matches []Match, tz string, includeInfoLinks bool) (result string) {
	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].BeginAt.Before(matches[j].BeginAt)
	})

	for _, m := range matches {
		parsedMatch, err := m.Parse(tz, includeInfoLinks)
		if err != nil {
			continue
		}
		result += parsedMatch
	}
	return result
}

func (m Match) Parse(tz string, includeInfoLinks bool) (parsed string, err error) {
	convertedTime, err := TimeIn(m.BeginAt, tz)
	if err != nil {
		return "", err
	}

	parsed += fmt.Sprintf(`*%s* \[BO%d]`, m.Name, m.NumberOfGames)
	parsed += "\n"
	parsed += convertedTime.Format(time.RFC822)
	parsed += "\n"
	if (m.Status == "running" || m.Status == "finished") && len(m.Results) == 2 {
		var team1Name, team2Name string
		var team1Score, team2Score int

		team1Score = m.Results[0].Score
		team2Score = m.Results[1].Score
		for _, opp := range m.Opponents {
			if opp.Opponent.ID == m.Results[0].TeamID {
				team1Name = opp.Opponent.Name
			}
			if opp.Opponent.ID == m.Results[1].TeamID {
				team2Name = opp.Opponent.Name
			}
		}
		parsed += fmt.Sprintf(`_%s_ \[%d:%d] _%s_`, team1Name, team1Score, team2Score, team2Name)
		parsed += "\n"
	}

	if m.Status != "finished" {
		for _, stream := range m.StreamsList {
			parsed += fmt.Sprintf("Twitch %s: <%s>\n", strings.ToUpper(stream.Language), stream.RawURL)
		}
	}

	if m.Tournament.Slug != "" {
		parsed += fmt.Sprintf("_%s_\n", strings.ReplaceAll(m.Tournament.Slug, "_", "-"))
	}
	if includeInfoLinks {
		if m.SerieID != 0 {
			parsed += fmt.Sprintf("Serie info: /serie@%d\n", m.SerieID)
		}
		if m.TournamentID != 0 {
			parsed += fmt.Sprintf("Tournament info: /tournament@%d\n", m.TournamentID)
		}
	}
	parsed += "\n"

	return parsed, nil
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
