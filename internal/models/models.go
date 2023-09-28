package models

import "time"

type Link struct {
	ID             string    `db:"id"`
	ShortSuffix    string    `db:"short_suffix"`
	Url            string    `db:"url"`
	Clicks         uint      `db:"clicks"`
	ExpirationDate time.Time `db:"expiration_date"`
}
