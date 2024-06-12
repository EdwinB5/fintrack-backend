package internal

import "time"

type Income struct {
	ID           string    `json:"id"`
	Concept      string    `json:"concept"`
	Account      string    `json:"account"`
	Total        float64   `json:"total"`
	Date         time.Time `json:"date"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`
}
