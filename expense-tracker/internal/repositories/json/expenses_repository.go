package json

import (
	"context"

	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"
)

type expensesRepository struct {
}

func NewExpensesRepository() domain.ExpensesRepository {
	return &expensesRepository{}
}

func (r *expensesRepository) GetAll(ctx context.Context) ([]domain.Expense, error) {
	return []domain.Expense{}, nil
}

func (r *expensesRepository) SaveAll(ctx context.Context, expenses []domain.Expense) error {
	return nil
}
