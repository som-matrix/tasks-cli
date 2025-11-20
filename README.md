# Tasks CLI

A simple command-line interface (CLI) task manager written in Go. This tool allows you to manage your daily tasks directly from your terminal.

## Features

- **Add Tasks**: Quickly add new tasks to your list.
- **List Tasks**: View all your current tasks.
- **Mark as Done**: Mark tasks as completed.
- **Delete Tasks**: Remove tasks from your list.
- **Persistent Storage**: Tasks are saved to a `tasks.json` file in the current directory.

## Installation

Ensure you have [Go](https://go.dev/dl/) installed on your machine.

1.  Clone the repository:
    ```bash
    git clone <repository-url>
    cd tasks
    ```

2.  Build the project:
    ```bash
    go build -o tasks cmd/tasks/main.go
    ```

3.  (Optional) Move the binary to your PATH to run it from anywhere.

## Usage

You can run the tool using the built binary or directly with `go run`.

### Add a Task
Add a new task by providing a description.

```bash
./tasks add "Buy groceries"
# or
go run cmd/tasks/main.go add "Buy groceries"
```

### List Tasks
View all tasks with their IDs and status.

```bash
./tasks list
```

### 5. Interactive Mode (TUI)
Launch an interactive terminal interface to manage tasks.
```bash
./tasks interactive
# OR
./tasks i
```
**Controls:**
- **Up/Down Arrow**: Navigate tasks
- **Enter**: Toggle "Done" status
- **q**: Quit

### 6. Help
Show usage information.
```bash
./tasks help
```

### Mark a Task as Done
Mark a task as completed using its ID (obtained from the `list` command).

```bash
./tasks done 1
```

### Delete a Task
Delete a task permanently using its ID.

```bash
./tasks delete 1
```

## Project Structure

- `cmd/tasks/main.go`: Entry point of the application, handles command-line arguments.
- `pkg/tasks/tasks.go`: Contains the core logic for task management (add, list, done, delete) and file operations.
- `tasks.json`: The file where tasks are stored (created automatically).
