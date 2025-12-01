package domain

import "time"

type Expense struct {
	Id          int64     `json:"id"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}
