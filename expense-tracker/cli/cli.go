package cli

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/domain"
	"github.com/Taras-Rm/technical-practice/expense-tracker/internal/services"
	"github.com/jedib0t/go-pretty/v6/table"
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

	switch command {
	case domain.COMMAND_ADD:
		err = cli.handleAdd(ctx, args[2:])
	case domain.COMMAND_LIST:
		err = cli.handleList(ctx)
	case domain.COMMAND_SUMMARY:
		err = cli.handleSummary(ctx, args[2:])
	case domain.COMMAND_DELETE:
		err = cli.handleDelete(ctx, args[2:])
	case domain.COMMAND_UPDATE:
		err = cli.handleUpdate(ctx, args[2:])
	default:
		fmt.Println("not supported command")
	}

	if err != nil {
		return err
	}

	return nil
}

func (cli *CLI) getCommand(args []string) (domain.CLICommand, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("wrong command")
	}

	return domain.CLICommand(args[1]), nil
}

func (cli *CLI) handleAdd(ctx context.Context, args []string) error {
	if len(args) != 6 {
		return fmt.Errorf("wrong add parameters")
	}

	if args[0] != "--description" {
		return fmt.Errorf("no description param")
	}

	if args[2] != "--amount" {
		return fmt.Errorf("no amount param")
	}

	if args[4] != "--category" {
		return fmt.Errorf("no category param")
	}

	amount, err := strconv.Atoi(args[3])
	if err != nil {
		return err
	}

	expense, err := cli.expensesService.AddExpense(ctx, args[1], int64(amount), args[5])
	if err != nil {
		return err
	}

	fmt.Printf("Expense added successfully (ID: %s)", expense.Id)

	return nil
}

func (cli *CLI) handleList(ctx context.Context) error {
	expenses, err := cli.expensesService.GetAllExpenses(ctx, domain.GetAllExpensesFilter{})
	if err != nil {
		return err
	}

	cli.printExpensesList(expenses)

	return nil
}

func (cli *CLI) printExpensesList(expenses []domain.Expense) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Date", "Description", "Amount", "Category"})

	var rows []table.Row

	for _, e := range expenses {
		rows = append(rows, table.Row{e.Id, e.Date.Format("2006-01-02"), e.Description, fmt.Sprintf("$%d", e.Amount), e.Category})
	}

	t.AppendRows(rows)
	t.AppendSeparator()
	t.Render()
}

func (cli *CLI) handleSummary(ctx context.Context, args []string) error {
	if len(args) != 0 && len(args) != 2 {
		return fmt.Errorf("wrong summary parameters")
	}

	var err error
	var monthFilter int

	if len(args) == 2 {
		if args[0] != "--month" {
			return fmt.Errorf("no month param")
		}

		monthFilter, err = strconv.Atoi(args[1])
		if err != nil {
			return err
		}
	}

	filter := domain.GetAllExpensesFilter{
		Month: time.Month(monthFilter),
	}

	expenses, err := cli.expensesService.GetAllExpenses(ctx, filter)
	if err != nil {
		return err
	}

	var totalExpenses int64

	for _, e := range expenses {
		totalExpenses += e.Amount
	}

	fmt.Printf("Total expenses: $%d", totalExpenses)

	return nil
}

func (cli *CLI) handleDelete(ctx context.Context, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("wrong delete parameters")
	}

	if args[0] != "--id" {
		return fmt.Errorf("no id param")
	}

	err := cli.expensesService.DeleteExpense(ctx, args[1])
	if err != nil {
		return err
	}

	fmt.Println("Expense deleted successfully")

	return nil
}

func (cli *CLI) handleUpdate(ctx context.Context, args []string) error {
	if len(args) != 8 {
		return fmt.Errorf("wrong update parameters")
	}

	if args[0] != "--id" {
		return fmt.Errorf("no description param")
	}

	if args[2] != "--description" {
		return fmt.Errorf("no description param")
	}

	if args[4] != "--amount" {
		return fmt.Errorf("no amount param")
	}

	if args[6] != "--category" {
		return fmt.Errorf("no category param")
	}

	amount, err := strconv.Atoi(args[5])
	if err != nil {
		return err
	}

	expense, err := cli.expensesService.UpdateExpense(ctx, args[1], args[3], int64(amount), args[7])
	if err != nil {
		return err
	}

	fmt.Printf("Expense updated successfully (ID: %s)", expense.Id)

	return nil
}
