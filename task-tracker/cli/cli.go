package cli

import (
	"fmt"
	"strconv"

	"github.com/Taras-Rm/technical-practice/task-tracker/db"
)

type CLI struct {
	taskStore db.TaskStore
}

func NewCLI(taskStore db.TaskStore) *CLI {
	return &CLI{
		taskStore,
	}
}

func (c *CLI) ParseCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("invalid command")
	}

	args = args[1:]

	switch args[0] {
	case "add":
		return c.handleAdd(args)
	case "delete":
		return c.handleDelete(args)
	case "update":
		return c.handleUpdate(args)
	default:
		fmt.Println("Wrong command.")
	}

	return nil
}

func (c *CLI) handleAdd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("description required")
	}

	description := args[0]

	task, err := c.taskStore.Add(description)
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)", task.Id)

	return nil
}

func (c *CLI) handleDelete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("id required")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	err = c.taskStore.Delete(id)
	if err != nil {
		return err
	}

	fmt.Printf("Task added deleted (ID: %d)", id)

	return nil
}

func (c *CLI) handleUpdate(args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("wrong command")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	description := args[1]

	err = c.taskStore.Update(id, description)
	if err != nil {
		return err
	}

	fmt.Printf("Task updated (ID: %d)", id)

	return nil
}
