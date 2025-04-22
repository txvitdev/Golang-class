package entities

import (
	"time"

	"github.com/lib/pq"
)

type Client struct {
	ID           string         `db:"id" json:"id"`
	Name         string         `db:"name" json:"name"`
	ClientID     string         `db:"client_id" json:"client_id"`
	ClientSecret string         `db:"client_secret" json:"client_secret"`
	RedirectURIs pq.StringArray `db:"redirect_uris" json:"redirect_uris"`
	CreatedAt    time.Time      `db:"created_at" json:"created_at"`
}
