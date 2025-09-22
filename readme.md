# Task Tracker CLI

Task Tracker CLI is a lightweight command-line application for capturing, updating, and tracking personal tasks. It implements the roadmap.sh Task Tracker project and runs on Windows, macOS, and Linux.

## Features
- Add new tasks with the default `todo` status.
- List all tasks or filter them by status (`todo`, `in-progress`, `done`).
- Update task descriptions, move tasks to `in-progress` or `done`, and delete tasks you no longer need.
- Store data locally in a JSON file so it is easy to back up or share.

## Requirements
- [Go](https://go.dev/dl/) version 1.25 or later (per the `go.mod` file).
- Git (optional but recommended for cloning the repository).

## Installation
1. Clone this repository and move into the project directory:
   ```bash
   git clone https://github.com/albantani17/Task-Tracker.git
   cd Task-Tracker
   ```
2. Build the CLI binary:
   ```bash
   go build -o task-cli ./cmd/task-tracker
   ```
   When the compilation completes, the `task-cli` binary (`task-cli.exe` on Windows) will be available in the project directory.

   Alternative: install directly into your `$GOBIN` using `go install` (Go 1.17+):
   ```bash
   go install github.com/albantani17/Task-Tracker/cmd/task-tracker@latest
   ```
   After installing, ensure `$GOBIN` or `$GOPATH/bin` is on your `PATH` so that `task-cli` can be invoked from anywhere.

## Usage
The general command structure is:
```bash
task-cli <command> [args]
```
Available commands:

| Command | Syntax | Description |
| --- | --- | --- |
| Help | `task-cli help` | Show a summary of the available commands. |
| Add | `task-cli add "Task Name"` | Add a new task with the `todo` status. |
| List | `task-cli list` | Display every stored task. |
| Filtered list | `task-cli list <status>` | Show only tasks matching a status (`todo`, `in-progress`, `done`). |
| Mark in-progress | `task-cli mark-in-progress <task_id>` | Change a task status to `in-progress`. |
| Mark done | `task-cli mark-done <task_id>` | Change a task status to `done`. |
| Update description | `task-cli update <task_id> "New Description"` | Replace the task description. |
| Delete | `task-cli delete <task_id>` | Remove a task by its ID. |

### Example session
```bash
# Add a new task
task-cli add "Learn Go"

# Inspect every task
task-cli list
# The output includes ID, description, color-coded status, created timestamp, and updated timestamp.

# Move the task to in-progress
task-cli mark-in-progress Tp9sLm4Q

# Update the description
task-cli update Tp9sLm4Q "Learn advanced Go"

# Mark it as done
task-cli mark-done Tp9sLm4Q

# Delete the task
task-cli delete Tp9sLm4Q
```
> Replace `Tp9sLm4Q` with the ID printed in your output.

## Data storage
- Tasks are stored in `storage/tasks.json`.
- If the file does not exist, the CLI creates it the first time you add a task.
- You can move or back up this file to preserve your task history.

## Development
Run the application without building a binary:
```bash
go run ./cmd/task-tracker <command> [args]
```

After modifying the source, run `go build` or `go run` again to validate that your changes behave as expected.
