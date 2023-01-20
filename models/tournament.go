package models

import "time"

type Tournament struct {
	BeginAt        time.Time   `json:"begin_at"`
	EndAt          interface{} `json:"end_at"`
	ExpectedRoster []struct {
		Players []struct {
			BirthYear   int         `json:"birth_year"`
			Birthday    string      `json:"birthday"`
			FirstName   string      `json:"first_name"`
			Hometown    string      `json:"hometown"`
			ID          int         `json:"id"`
			ImageURL    string      `json:"image_url"`
			LastName    string      `json:"last_name"`
			Name        string      `json:"name"`
			Nationality string      `json:"nationality"`
			Role        interface{} `json:"role"`
			Slug        string      `json:"slug"`
		} `json:"players"`
		Team struct {
			Acronym    interface{} `json:"acronym"`
			ID         int         `json:"id"`
			ImageURL   string      `json:"image_url"`
			Location   string      `json:"location"`
			ModifiedAt time.Time   `json:"modified_at"`
			Name       string      `json:"name"`
			Slug       string      `json:"slug"`
		} `json:"team"`
	} `json:"expected_roster"`
	ID     int `json:"id"`
	League struct {
		ID         int         `json:"id"`
		ImageURL   string      `json:"image_url"`
		ModifiedAt time.Time   `json:"modified_at"`
		Name       string      `json:"name"`
		Slug       string      `json:"slug"`
		URL        interface{} `json:"url"`
	} `json:"league"`
	LeagueID      int       `json:"league_id"`
	LiveSupported bool      `json:"live_supported"`
	Matches       []Match   `json:"matches"`
	ModifiedAt    time.Time `json:"modified_at"`
	Name          string    `json:"name"`
	Prizepool     string    `json:"prizepool"`
	Serie         struct {
		BeginAt     time.Time   `json:"begin_at"`
		Description interface{} `json:"description"`
		EndAt       interface{} `json:"end_at"`
		FullName    string      `json:"full_name"`
		ID          int         `json:"id"`
		LeagueID    int         `json:"league_id"`
		ModifiedAt  time.Time   `json:"modified_at"`
		Name        string      `json:"name"`
		Season      string      `json:"season"`
		Slug        string      `json:"slug"`
		WinnerID    interface{} `json:"winner_id"`
		WinnerType  interface{} `json:"winner_type"`
		Year        int         `json:"year"`
	} `json:"serie"`
	SerieID int    `json:"serie_id"`
	Slug    string `json:"slug"`
	Teams   []struct {
		Acronym    interface{} `json:"acronym"`
		ID         int         `json:"id"`
		ImageURL   string      `json:"image_url"`
		Location   string      `json:"location"`
		ModifiedAt time.Time   `json:"modified_at"`
		Name       string      `json:"name"`
		Slug       string      `json:"slug"`
	} `json:"teams"`
	Videogame struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"videogame"`
	WinnerID   interface{} `json:"winner_id"`
	WinnerType interface{} `json:"winner_type"`
}
