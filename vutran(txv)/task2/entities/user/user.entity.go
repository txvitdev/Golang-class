package entities

import "time"

type User struct {
	Id        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
}