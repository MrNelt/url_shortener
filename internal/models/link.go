package models

import "time"

type Link struct {
	SecretKey   string `db:"secret_key"`
	ShortSuffix string `db:"short_suffix"`
	Url         string `db:"url"`
	Clicks      uint   `db:"clicks"`
}

type LinkWithTTL struct {
	SecretKey      string    `db:"secret_key"`
	ShortSuffix    string    `db:"short_suffix"`
	Url            string    `db:"url"`
	ExpirationDate time.Time `db:"expiration_date"`
	Clicks         uint      `db:"clicks"`
}
