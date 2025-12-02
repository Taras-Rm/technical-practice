package cli

import (
	"context"
	"fmt"
	"strconv"

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

func (cli *CLI) Run(ctx context.Context, args []string) error {
	command, err := cli.getCommand(args)
	if err != nil {
		return err
	}

	var response string

	switch command {
	case domain.COMMAND_ADD:
		response, err = cli.handleAdd(ctx, args[2:])
	case domain.COMMAND_LIST:
		response, err = cli.handleList(ctx)
	case domain.COMMAND_SUMMARY:
		response, err = cli.handleSummary(ctx)
	case domain.COMMAND_DELETE:
		response, err = cli.handleDelete(ctx, args[2:])
	case domain.COMMAND_UPDATE:
		response, err = cli.handleUpdate(ctx, args[2:])
	default:
		response = "not supported command"
	}

	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}

func (cli *CLI) getCommand(args []string) (domain.CLICommand, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("wrong command")
	}

	return domain.CLICommand(args[1]), nil
}

func (cli *CLI) handleAdd(ctx context.Context, args []string) (string, error) {
	if len(args) != 4 {
		return "", fmt.Errorf("wrong add parameters")
	}

	if args[0] != "--description" {
		return "", fmt.Errorf("no description param")
	}

	if args[2] != "--amount" {
		return "", fmt.Errorf("no amount param")
	}

	amount, err := strconv.Atoi(args[3])
	if err != nil {
		return "", err
	}

	expense, err := cli.expensesService.AddExpense(ctx, args[1], int64(amount))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Expense added successfully (ID: %d)", expense.Id), nil
}

func (cli *CLI) handleList(ctx context.Context) (string, error) {
	expenses, err := cli.expensesService.GetAllExpenses(ctx)
	if err != nil {
		return "", err
	}

	var res string

	for _, e := range expenses {
		res += fmt.Sprintf("%d   %s  %s        %d\n", e.Id, e.Date.String(), e.Description, e.Amount)
	}

	return res, nil
}

func (cli *CLI) handleSummary(ctx context.Context) (string, error) {
	expenses, err := cli.expensesService.GetAllExpenses(ctx)
	if err != nil {
		return "", err
	}

	var totalExpenses int64

	for _, e := range expenses {
		totalExpenses += e.Amount
	}

	return fmt.Sprintf("Total expenses: $%d", totalExpenses), nil
}

func (cli *CLI) handleDelete(ctx context.Context, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("wrong delete parameters")
	}

	if args[0] != "--id" {
		return "", fmt.Errorf("no id param")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	err = cli.expensesService.DeleteExpense(ctx, int64(id))
	if err != nil {
		return "", err
	}

	return "Expense deleted successfully", nil
}

func (cli *CLI) handleUpdate(ctx context.Context, args []string) (string, error) {
	if len(args) != 6 {
		return "", fmt.Errorf("wrong update parameters")
	}

	if args[0] != "--id" {
		return "", fmt.Errorf("no description param")
	}

	if args[2] != "--description" {
		return "", fmt.Errorf("no description param")
	}

	if args[4] != "--amount" {
		return "", fmt.Errorf("no amount param")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return "", err
	}

	amount, err := strconv.Atoi(args[5])
	if err != nil {
		return "", err
	}

	expense, err := cli.expensesService.UpdateExpense(ctx, int64(id), args[3], int64(amount))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Expense updated successfully (ID: %d)", expense.Id), nil
}
