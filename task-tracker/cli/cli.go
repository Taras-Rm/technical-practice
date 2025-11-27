package cli

import (
	"fmt"

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

	switch args[0] {
	case "add":
		return c.handleAdd(args[1:])
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
