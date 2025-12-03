package domain

import "time"

type Expense struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
	Category    string    `json:"category"`
}

type GetAllExpensesFilter struct {
	Month time.Month
}
