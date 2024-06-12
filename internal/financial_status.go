package internal

import "time"

type FinancialStatus struct {
	TotalIncome  float64   `json:"total_income"`
	TotalExpense float64   `json:"total_expense"`
	Current      float64   `json:"current"`
	TotalLoaned  float64   `json:"total_loaned"`
	Date         time.Time `json:"date"`
	DateCreated  time.Time `json:"date_created"`
	DateModified time.Time `json:"date_modified"`
}
