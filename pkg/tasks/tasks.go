package tasks

import (
	"fmt"
	"os"
	"strings"
)

var fileName = "tasks.txt"

func Add(taskDescription string) {
	// TODO: Implement the logic to save the task to a file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating tasks file:", err)
		return
	}
	defer file.Close()
	if taskDescription == "" {
		fmt.Println("Task description is empty")
		return
	}
	_, err = file.WriteString(taskDescription + "\n")
	if err != nil {
		fmt.Println("Error writing to tasks file:", err)
		return
	}
	fmt.Println("Adding task:", taskDescription)
}

func List() {
	// TODO: Implement the logic to read the tasks from the file
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return
	}
	fmt.Println(string(content))
}

func MarkAsDone(taskID int) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return
	}

	// Split tasks into lines
	lines := strings.Split(string(content), "\n")

	// ignore if task number is invalid
	if taskID <= 0 || taskID > len(lines) {
		fmt.Println("Invalid task number")
		return
	}

	if strings.HasPrefix(lines[taskID-1], "✅ ") {
		fmt.Println("Task", taskID, "is already done")
		return
	}

	lines[taskID-1] = "✅ " + lines[taskID-1]
	newContent := strings.Join(lines, "\n")
	os.WriteFile(fileName, []byte(newContent), 0644)
	fmt.Println("Task", taskID, "marked as done")
}

func Delete(taskID int) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading tasks file:", err)
		return
	}

	// Split tasks into lines
	lines := strings.Split(string(content), "\n")

	// ignore if task number is invalid
	if taskID <= 0 || taskID > len(lines) {
		fmt.Println("Invalid task number")
		return
	}

	// delete the task
	lines = append(lines[:taskID-1], lines[taskID:]...)
	newContent := strings.Join(lines, "\n")
	os.WriteFile(fileName, []byte(newContent), 0644)
	fmt.Println("Task", taskID, "deleted")
}
