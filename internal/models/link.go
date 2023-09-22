package models

type Link struct {
	SecretKey   string `db:"secret_key"`
	ShortSuffix string `db:"short_suffix"`
	Url         string `db:"url"`
	Clicks      uint   `db:"clicks"`
}
