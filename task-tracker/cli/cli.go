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

	handleArgs := args[1:]

	switch args[0] {
	case "add":
		return c.handleAdd(handleArgs)
	case "delete":
		return c.handleDelete(handleArgs)
	case "update":
		return c.handleUpdate(handleArgs)
	case "list":
		return c.handleList(handleArgs)
	case "mark-in-progress":
		return c.handleSetInProgress(handleArgs)
	case "mark-done":
		return c.handleSetDone(handleArgs)
	default:
		fmt.Println("Wrong command.")
	}

	return nil
}

func (c *CLI) handleAdd(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("description required")
	}

	task, err := c.taskStore.Add(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)", task.Id)

	return nil
}

func (c *CLI) handleDelete(args []string) error {
	if len(args) < 1 {
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
	if len(args) < 2 {
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

func (c *CLI) handleList(args []string) error {
	status := ""

	if len(args) > 0 {
		status = args[0]
	}

	tasks, err := c.taskStore.GetAll(status)
	if err != nil {
		return err
	}

	fmt.Printf("Tasks: %+v", tasks)

	return nil
}

func (c *CLI) handleSetInProgress(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("id required")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	err = c.taskStore.SetStatus(id, "in-progress")
	if err != nil {
		return err
	}

	fmt.Printf("Updated status: %d", id)

	return nil
}

func (c *CLI) handleSetDone(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("id required")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	err = c.taskStore.SetStatus(id, "done")
	if err != nil {
		return err
	}

	fmt.Printf("Updated status: %d", id)

	return nil
}
