package models

import "time"

type Serie struct {
	BeginAt     time.Time   `json:"begin_at"`
	Description interface{} `json:"description"`
	EndAt       time.Time   `json:"end_at"`
	FullName    string      `json:"full_name"`
	ID          int         `json:"id"`
	League      struct {
		ID         int         `json:"id"`
		ImageURL   string      `json:"image_url"`
		ModifiedAt time.Time   `json:"modified_at"`
		Name       string      `json:"name"`
		Slug       string      `json:"slug"`
		URL        interface{} `json:"url"`
	} `json:"league"`
	LeagueID    int         `json:"league_id"`
	ModifiedAt  time.Time   `json:"modified_at"`
	Name        interface{} `json:"name"`
	Season      string      `json:"season"`
	Slug        string      `json:"slug"`
	Tournaments []struct {
		BeginAt       time.Time   `json:"begin_at"`
		EndAt         time.Time   `json:"end_at"`
		ID            int         `json:"id"`
		LeagueID      int         `json:"league_id"`
		LiveSupported bool        `json:"live_supported"`
		ModifiedAt    time.Time   `json:"modified_at"`
		Name          string      `json:"name"`
		Prizepool     interface{} `json:"prizepool"`
		SerieID       int         `json:"serie_id"`
		Slug          string      `json:"slug"`
		WinnerID      int         `json:"winner_id"`
		WinnerType    string      `json:"winner_type"`
	} `json:"tournaments"`
	Videogame struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"videogame"`
	WinnerID   int    `json:"winner_id"`
	WinnerType string `json:"winner_type"`
	Year       int    `json:"year"`
}
