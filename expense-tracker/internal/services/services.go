package services

import "github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"

type ExpensesService interface {
	AddExpense(description string, amount int64) (*domain.Expense, error)
	DeleteExpense(id int64) error
	UpdateExpense(id int64, description string, amount int64) (*domain.Expense, error)
	GetAllExpenses() ([]domain.Expense, error)
}
