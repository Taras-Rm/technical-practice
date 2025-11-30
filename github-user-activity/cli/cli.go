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

	fmt.Println(publicEvents)
	fmt.Println("finish")

	return nil
}

func (c *CLI) getUsername(args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("not enough arguments")
	}

	return args[1], nil
}
