package main

import (
	"fmt"

	"github.com/radityajay/task-tracker-cli/cli"
	"github.com/radityajay/task-tracker-cli/task"
)

func main() {
	repo := task.NewRepository("task.json")
	uc := task.NewUseCase(repo)
	err := cli.Run(*uc)
	if err != nil {
		fmt.Println(err)
	}
}