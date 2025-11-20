package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"tasks/pkg/tasks"
	"tasks/pkg/tui"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(os.Args[2:])
	case "list":
		handleList()
	case "done":
		handleDone(os.Args[2:])
	case "delete":
		handleDelete(os.Args[2:])
	case "interactive", "i":
		tui.Start()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: tasks <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  add <description>   Add a new task (no quotes needed)")
	fmt.Println("    Flags:")
	fmt.Println("      -priority, -p   Set priority (High, Medium, Low)")
	fmt.Println("  list                List all tasks")
	fmt.Println("  done <id>           Mark a task as done")
	fmt.Println("  delete <id>         Delete a task")
	fmt.Println("  interactive, i      Start interactive mode")
	fmt.Println("\nExamples:")
	fmt.Println("  tasks add Buy milk")
	fmt.Println("  tasks add Buy milk -priority High")
	fmt.Println("  tasks list")
	fmt.Println("  tasks done 1")
	fmt.Println("  tasks interactive")
}

func handleAdd(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a task description.")
		os.Exit(1)
	}

	var descriptionParts []string
	priority := "Low"

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "-priority" || arg == "-p" {
			if i+1 < len(args) {
				priority = args[i+1]
				i++ // Skip the next argument since it's the priority value
			} else {
				fmt.Println("Error: Please provide a value for priority.")
				os.Exit(1)
			}
		} else {
			descriptionParts = append(descriptionParts, arg)
		}
	}

	if len(descriptionParts) == 0 {
		fmt.Println("Error: Please provide a task description.")
		os.Exit(1)
	}

	taskDescription := strings.Join(descriptionParts, " ")
	tasks.Add(taskDescription, priority)
}

func handleList() {
	tasks.List()
}

func handleDone(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a task ID.")
		os.Exit(1)
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: Invalid task ID. Please provide a number.")
		os.Exit(1)
	}
	tasks.MarkAsDone(taskID)
}

func handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a task ID.")
		os.Exit(1)
	}
	taskID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Error: Invalid task ID. Please provide a number.")
		os.Exit(1)
	}
	tasks.Delete(taskID)
}
