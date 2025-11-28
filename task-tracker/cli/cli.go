package cli

import (
	"fmt"
	"strconv"

	"github.com/Taras-Rm/technical-practice/task-tracker/db"
	"github.com/Taras-Rm/technical-practice/task-tracker/domain"
)

const (
	ADD_CMD              = "add"
	DELETE_CMD           = "delete"
	UPDATE_CMD           = "update"
	LIST_CMD             = "list"
	MARK_IN_PROGRESS_CMD = "mark_in_progress"
	MARK_DONE_CMD        = "mark_done"
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
	case ADD_CMD:
		return c.handleAdd(handleArgs)
	case DELETE_CMD:
		return c.handleDelete(handleArgs)
	case UPDATE_CMD:
		return c.handleUpdate(handleArgs)
	case LIST_CMD:
		return c.handleList(handleArgs)
	case MARK_IN_PROGRESS_CMD:
		return c.handleSetInProgress(handleArgs)
	case MARK_DONE_CMD:
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

	err = c.taskStore.SetStatus(id, string(domain.StatusInProgress))
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

	err = c.taskStore.SetStatus(id, string(domain.StatusDone))
	if err != nil {
		return err
	}

	fmt.Printf("Updated status: %d", id)

	return nil
}
