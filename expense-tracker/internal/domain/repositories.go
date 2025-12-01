package domain

import "context"

type ExpensesRepository interface {
	GetAll(ctx context.Context) ([]Expense, error)
	SaveAll(ctx context.Context, expenses []Expense) error
}
