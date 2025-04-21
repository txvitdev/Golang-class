package entities

import "time"

type User struct {
	Id        int64  `db:"id" json:"id" form:"id"`
	Email     string `db:"email" json:"email" form:"email"`
	Password  string `db:"password" json:"password" form:"password"`
	CreatedAt time.Time  `db:"created_at" json:"created_at" form:"created_at"`
}