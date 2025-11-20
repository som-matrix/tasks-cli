package tui

import (
	"fmt"
	"os"
	"tasks/pkg/tasks"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type model struct {
	tasks  []tasks.Task
	cursor int
	err    error
}

func initialModel() model {
	t, err := tasks.LoadTasks()
	if err != nil {
		return model{err: err}
	}
	return model{tasks: t}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}

		case "enter", " ":
			selectedTask := &m.tasks[m.cursor]
			// Toggle done status
			selectedTask.Done = !selectedTask.Done
			// Save changes
			tasks.SaveTasks(m.tasks)
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if len(m.tasks) == 0 {
		return "No tasks found.\nPress 'q' to quit.\n"
	}

	s := "Interactive Task List\n\n"

	for i, task := range m.tasks {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		checked := " " // not done
		if task.Done {
			checked = "x" // done
		}

		status := "⏳"
		if task.Done {
			status = "✅"
		}

		priority := task.Priority
		if priority == "" {
			priority = "Low"
		}

		priorityColor := color.New(color.FgBlue).SprintFunc()
		switch priority {
		case "High":
			priorityColor = color.New(color.FgRed).SprintFunc()
		case "Medium":
			priorityColor = color.New(color.FgYellow).SprintFunc()
		}

		line := fmt.Sprintf("%s [%s] %s (%s) - %s", cursor, checked, task.Description, priorityColor(priority), status)

		if m.cursor == i {
			line = color.New(color.Bold, color.FgCyan).Sprint(line)
		}

		s += line + "\n"
	}

	s += "\nPress q to quit.\n"
	return s
}

func Start() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
