package main

import (
	"fmt"
	"os"

	"github.com/Taras-Rm/technical-practice/task-tracker/db"
)

const tasksStoreFileName = "./tasks.json"

func main() {
	args := os.Args
	fmt.Println("Hello world", args)

	tasksStore := db.InitTaskStore(tasksStoreFileName)

	err := tasksStore.Update(1, "Fishing")
	if err != nil {
		fmt.Printf("Error happend: %s", err.Error())
	}

	// fmt.Println(tasks)
}
