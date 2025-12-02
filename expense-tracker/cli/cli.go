package cli

import (
	"fmt"

	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"
	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/services"
)

type CLI struct {
	expensesService services.ExpensesService
}

func NewCLI(expensesService services.ExpensesService) *CLI {
	return &CLI{
		expensesService,
	}
}

func (cli *CLI) Run(args []string) error {
	_, err := cli.getCommand(args)
	if err != nil {
		return err
	}

	// 2. handle specific command swith

	return nil
}

func (cli *CLI) getCommand(args []string) (domain.CLICommand, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("wrong command")
	}

	return domain.CLICommand(args[1]), nil
}
