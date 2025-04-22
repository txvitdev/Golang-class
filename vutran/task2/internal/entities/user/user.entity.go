package entities

import "time"

type User struct {
	Id              string    `db:"id" json:"id" form:"id"`
	Email           string    `db:"email" json:"email" form:"email"`
	Password        string    `db:"password" json:"password" form:"password"`
	FullName        string    `db:"full_name" json:"full_name" form:"full_name"`
	IsEmailVerified bool      `db:"is_email_verified" json:"is_email_verified" form:"is_email_verified"`
	CreatedAt       time.Time `db:"created_at" json:"created_at" form:"created_at"`
}
