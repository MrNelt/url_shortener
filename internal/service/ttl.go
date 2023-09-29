package service

import (
	"time"
	errorsApi "url_shortener/internal/errors"
)

func _TTLDTOToDate(TTLCount uint, TTLUnit string) (time.Time, error) {
	date := time.Now()
	switch TTLUnit {
	case "DAYS":
		date = date.Add(time.Hour * 24 * time.Duration(TTLCount))
	case "HOURS":
		date = date.Add(time.Hour * time.Duration(TTLCount))
	case "MINUTES":
		date = date.Add(time.Minute * time.Duration(TTLCount))
	case "SECONDS":
		date = date.Add(time.Second * time.Duration(TTLCount))
	default:
		return time.Time{}, errorsApi.ErrTTL
	}
	return date, nil
}
