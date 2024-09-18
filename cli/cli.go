package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/radityajay/task-tracker-cli/task"
)

func Run(uc task.UseCase) error {
	command := make([]string, 3)
	copy(command, os.Args[1:])

	switch command[0] {
	case "add":
		// TODO: Implement add task logic
		fmt.Println("Adding a new task...")
		return uc.Add(&task.Task{
			ID:          0,
			Description: command[1],
			Status:      task.STATUS_TODO,
		})
	case "update":
		// TODO: Implement update task logic
		fmt.Println("Updating a task...")
		id, err := strconv.Atoi(command[1])
		if err != nil {
			return err
		}
		return uc.Update(uint(id), command[2])
	case "list":
		// TODO: Implement get task logic
		fmt.Println("Getting tasks...")
		fmt.Println("Command:", len(command))
		var tasks []*task.Task
		var err error
		if len(command) == 2 {
			fmt.Println("Getting all tasks...")
			tasks, err = uc.List(task.Status(command[1]))
			if err != nil {
				return err
			}
		}

		if len(command) == 3 {
			fmt.Println("Getting tasks by status...")
			tasks, err = uc.GetListByStatus(task.Status(command[1]))
			if err != nil {
				return err
			}
		}

		for no, task := range tasks {
			fmt.Printf("%d. ID: %d, Status: %s, Description: %s\n", no+1, task.ID, task.Status, task.Description)
		}
	case "delete":
		fmt.Println("Deleting a task...")
		id, err := strconv.Atoi(command[1])
		if err != nil {
			return err
		}
		return uc.Delete(uint(id))
	case "mark-in-progress":
		fmt.Println("Marking a task in progress...")
		id, err := strconv.Atoi(command[1])
		if err != nil {
			return err
		}
		return uc.UpdateStatus(uint(id), task.STATUS_IN_PROGRESS)
	case "mark-done":
		fmt.Println("Marking a task done...")
		id, err := strconv.Atoi(command[1])
		if err != nil {
			return err
		}
		return uc.UpdateStatus(uint(id), task.STATUS_DONE)
	default:
		fmt.Println("Invalid command.")
	}

	return nil
}
