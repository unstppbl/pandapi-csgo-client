package models

type User struct {
	ID        int64  `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`

	State string `json:"state" bson:"state"`
}
