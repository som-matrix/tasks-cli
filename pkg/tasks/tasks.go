package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
)

var fileName = "tasks.json"

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
	Priority    string `json:"priority"`
}

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if len(content) == 0 {
		return []Task{}, nil
	}

	err = json.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	content, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, content, 0644)
}

func Add(taskDescription string, priority string) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}

	if priority == "" {
		priority = "Low"
	}

	newTask := Task{
		ID:          id,
		Description: taskDescription,
		Done:        false,
		Priority:    priority,
	}

	tasks = append(tasks, newTask)

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Printf("Task added successfully (ID: %d, Priority: %s)\n", id, priority)
}

func List() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

	// Header
	header := color.New(color.FgCyan, color.Bold, color.Underline).SprintFunc()
	fmt.Fprintln(w, header("ID\tTask\tStatus\tPriority"))

	for _, task := range tasks {
		status := color.YellowString("⏳ Pending")
		if task.Done {
			status = color.GreenString("✅ Done")
		}

		priority := task.Priority
		if priority == "" {
			priority = "Low"
		}

		priorityColored := priority
		switch priority {
		case "High":
			priorityColored = color.RedString(priority)
		case "Medium":
			priorityColored = color.YellowString(priority)
		case "Low":
			priorityColored = color.BlueString(priority)
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, status, priorityColored)
	}
	w.Flush()
}

func MarkAsDone(taskID int) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	for i, task := range tasks {
		if task.ID == taskID {
			if tasks[i].Done {
				fmt.Printf("Task %d is already marked as done.\n", taskID)
				return
			}
			tasks[i].Done = true
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found.\n", taskID)
		return
	}

	err = SaveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task %d marked as done.\n", taskID)
}

func Delete(taskID int) {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newTasks := []Task{}
	found := false
	for _, task := range tasks {
		if task.ID == taskID {
			found = true
			continue
		}
		newTasks = append(newTasks, task)
	}

	if !found {
		fmt.Printf("Task with ID %d not found.\n", taskID)
		return
	}

	err = SaveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Printf("Task %d deleted.\n", taskID)
}
