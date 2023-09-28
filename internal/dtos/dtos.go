package dtos

import "time"

type LinkDTO struct {
	ShortSuffix string `json:"short_suffix"`
	Url         string `json:"url"`
	TTLCount    uint   `json:"ttl_count"`
	TTLUnit     string `json:"ttl_unit"`
}

type LinkInfoDto struct {
	ID             string    `json:"id"`
	ShortSuffix    string    `json:"short_suffix"`
	Url            string    `json:"url"`
	Clicks         uint      `json:"clicks"`
	ExpirationDate time.Time `json:"expiration_date"`
}
