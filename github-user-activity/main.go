package main

import (
	"fmt"
	"os"

	"github.com/Taras-Rm/technical-practice/github-user-activity/cli"
	githubapi "github.com/Taras-Rm/technical-practice/github-user-activity/github-api"
)

func main() {
	api := githubapi.NewGithubAPI()

	cli := cli.NewCLI(api)

	err := cli.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
