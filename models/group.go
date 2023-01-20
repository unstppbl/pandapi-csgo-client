package models

import "time"

type Group struct {
	ID        int64     `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	Username  string    `json:"username" bson:"username"`
	Members   []int64   `json:"members" bson:"members"`
	OwnerID   int64     `json:"owner_id" bson:"owner_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
