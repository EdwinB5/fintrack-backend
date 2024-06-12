package internal

import "time"

type Expense struct {
	ID           string    `json:"id"`
	PaidType     string    `json:"paid_type"`
	Category     string    `json:"category"`
	Total        string    `json:"total"`
	Account      string    `json:"account"`
	Date         time.Time `json:"date"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`
}
