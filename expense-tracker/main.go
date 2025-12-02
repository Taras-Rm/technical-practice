package main

import (
	"os"

	"github.com/Taras-Rm/technical-practice/expense-tracker/cli"
	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/repositories/json"
	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/services"
)

func main() {
	expensesRepository := json.NewExpensesRepository()

	expensesService := services.NewExpensesService(expensesRepository)

	cli := cli.NewCLI(expensesService)

	err := cli.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
