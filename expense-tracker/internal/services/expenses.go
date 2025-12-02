package services

import (
	"context"
	"fmt"
	"time"

	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"
)

type expensesService struct {
	expensesRepo domain.ExpensesRepository
}

func NewExpensesService(expensesRepo domain.ExpensesRepository) ExpensesService {
	return &expensesService{
		expensesRepo,
	}
}

func (s *expensesService) AddExpense(ctx context.Context, description string, amount int64) (*domain.Expense, error) {
	expenses, err := s.expensesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	newExpense := domain.Expense{
		Id:          2,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}

	expenses = append(expenses, newExpense)

	err = s.expensesRepo.SaveAll(ctx, expenses)
	if err != nil {
		return nil, err
	}

	return &newExpense, nil
}

func (s *expensesService) DeleteExpense(ctx context.Context, id int64) error {
	expenses, err := s.expensesRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	var filteredExpenses []domain.Expense

	for _, expense := range expenses {
		if expense.Id != id {
			filteredExpenses = append(filteredExpenses, expense)
		}
	}

	if len(expenses) == len(filteredExpenses) {
		return fmt.Errorf("expense with id: %d not found", id)
	}

	err = s.expensesRepo.SaveAll(ctx, filteredExpenses)
	if err != nil {
		return err
	}

	return nil
}

func (s *expensesService) UpdateExpense(ctx context.Context, id int64, description string, amount int64) (*domain.Expense, error) {
	expenses, err := s.expensesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var filteredExpenses []domain.Expense

	var updatedExpense domain.Expense

	for _, expense := range expenses {
		if expense.Id == id {
			updatedExpense = domain.Expense{
				Id:          expense.Id,
				Description: description,
				Amount:      amount,
				Date:        expense.Date,
			}

			filteredExpenses = append(filteredExpenses, updatedExpense)
		}
	}

	err = s.expensesRepo.SaveAll(ctx, filteredExpenses)
	if err != nil {
		return nil, err
	}

	return &updatedExpense, nil
}

func (s *expensesService) GetAllExpenses(ctx context.Context, filter domain.GetAllExpensesFilter) ([]domain.Expense, error) {
	expenses, err := s.expensesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return s.filterExpenses(expenses, filter), nil
}

func (s *expensesService) filterExpenses(expenses []domain.Expense, filter domain.GetAllExpensesFilter) []domain.Expense {
	if filter.Month == 0 {
		return expenses
	}

	var filteredExpenses []domain.Expense

	if filter.Month > 0 {
		for _, expense := range expenses {
			if expense.Date.Month() == time.Month(filter.Month) {
				filteredExpenses = append(filteredExpenses, expense)
			}
		}
	}

	return filteredExpenses
}
