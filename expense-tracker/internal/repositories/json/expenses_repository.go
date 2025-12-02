package json

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"
)

const FileModePermit = 0644

type expensesRepository struct {
	fileName string
}

func NewExpensesRepository(fileName string) domain.ExpensesRepository {
	return &expensesRepository{
		fileName,
	}
}

func (r *expensesRepository) GetAll(ctx context.Context) ([]domain.Expense, error) {
	data, err := os.ReadFile(r.fileName)
	if err != nil {
		return nil, err
	}

	var storage domain.Storage

	err = json.Unmarshal(data, &storage)
	if err != nil {
		return nil, err
	}

	return storage.Expenses, nil
}

func (r *expensesRepository) SaveAll(ctx context.Context, expenses []domain.Expense) error {
	updatedStorage := domain.Storage{
		Expenses: expenses,
	}

	byteJson, err := json.Marshal(updatedStorage)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.fileName, byteJson, FileModePermit)
	if err != nil {
		return err
	}

	return nil
}
