package main

import (
	"fmt"
	"os"

	commCli "github.com/Taras-Rm/technical-practice/task-tracker/cli"
	"github.com/Taras-Rm/technical-practice/task-tracker/db"
)

const tasksStoreFileName = "./tasks.json"

func main() {
	args := os.Args

	tasksStore := db.InitTaskStore(tasksStoreFileName)

	cli := commCli.NewCLI(tasksStore)

	err := cli.ParseCommand(args[1:])

	if err != nil {
		fmt.Println(err)
	}
}
