package services

import (
	"context"

	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"
)

type ExpensesService interface {
	AddExpense(ctx context.Context, description string, amount int64) (*domain.Expense, error)
	DeleteExpense(ctx context.Context, id int64) error
	UpdateExpense(ctx context.Context, id int64, description string, amount int64) (*domain.Expense, error)
	GetAllExpenses(ctx context.Context) ([]domain.Expense, error)
}
