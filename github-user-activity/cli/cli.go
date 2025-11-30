package cli

import (
	"fmt"

	githubapi "github.com/Taras-Rm/technical-practice/github-user-activity/github-api"
)

type CLI struct {
	githubAPI githubapi.IGithubAPI
}

func NewCLI(githubAPI githubapi.IGithubAPI) *CLI {
	return &CLI{
		githubAPI,
	}
}

func (c *CLI) Run(args []string) error {
	username, err := c.getUsername(args)
	if err != nil {
		return err
	}

	publicEvents, err := c.githubAPI.GetUserActivity(username)
	if err != nil {
		return err
	}

	printResult := c.formatEventsForPrint(publicEvents)

	fmt.Println(printResult)

	return nil
}

func (c *CLI) getUsername(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough arguments")
	}

	return args[1], nil
}

func (c *CLI) formatEventsForPrint(events []githubapi.PublicEvent) string {
	print := ""

	for idx, event := range events {
		switch event.Type {
		case githubapi.PushEventType:
			print += fmt.Sprintf("Pushed commit to %s", event.Repo.Name)
		case githubapi.CreateEventType:
			print += fmt.Sprintf("Created new repository %s", event.Repo.Name)
		case githubapi.DeleteEventType:
			print += fmt.Sprintf("Deleted repository %s", event.Repo.Name)
		default:
			print += "unknown event type"
		}

		if idx < len(events)-1 {
			print += "\n"
		}
	}

	return print
}
