package domain

import (
	"time"
)

type Account struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Source    string    `bson:"source"`
	UserID    string    `bson:"user_id"`
	UserName  string    `bson:"user_name"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
