package models

import "time"

type Team struct {
	Acronym          string `json:"acronym"`
	CurrentVideogame struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"current_videogame"`
	ID         int       `json:"id"`
	ImageURL   string    `json:"image_url"`
	Location   string    `json:"location"`
	ModifiedAt time.Time `json:"modified_at"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	distance   int
}
