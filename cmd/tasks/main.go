package main

import (
	"fmt"
	"os"
	"strconv"
	"tasks/pkg/tasks"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tasks <command> [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
			return
		}
		taskDescription := os.Args[2]
		tasks.Add(taskDescription)

	case "list":
		if len(os.Args) > 2 {
			fmt.Println("Invalid number of arguments.")
			return
		}
		tasks.List()

	case "done":
		if len(os.Args) > 3 {
			fmt.Println("Invalid number of arguments.")
			return
		}
		taskIDStr := os.Args[2]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a number.")
			return
		}
		tasks.MarkAsDone(taskID)

	case "delete":
		if len(os.Args) > 3 {
			fmt.Println("Invalid number of arguments.")
			return
		}
		taskIDStr := os.Args[2]
		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			fmt.Println("Invalid task ID. Please provide a number.")
			return
		}
		tasks.Delete(taskID)

	default:
		fmt.Println("Unknown command:", command)
	}
}
