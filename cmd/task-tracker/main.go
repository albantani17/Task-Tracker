package main

import (
	"fmt"
	"github.com/albantani17/Task-Tracker/internal/task"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli <command> [args...]")
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <task name>")
			return
		}
		title := os.Args[2]
		task.AddTask(title)
		return
	case "list":
		if len(os.Args) > 2 {
			status := os.Args[2]
			task.FindTasksWithStatus(task.Status(status))
			return
		}
		task.FindTasks()
		return
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <task id>")
			return
		}
		id := os.Args[2]
		task.MarkTask(id, task.InProgress)
		return
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <task id>")
			return
		}
		id := os.Args[2]
		task.MarkTask(id, task.Done)
		return
	case "update":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli update <task id> <task name>")
			return
		}
		id := os.Args[2]
		title := os.Args[3]
		task.UpdateTask(id, title)
		return
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <task id>")
			return
		}
		id := os.Args[2]
		task.DeleteTask(id)
		return
	case "help":
		fmt.Println("Usage: task-cli list <status>")
		fmt.Println("Usage: task-cli add <task name>")
		fmt.Println("Usage: task-cli mark-in-progress <task id>")
		fmt.Println("Usage: task-cli mark-done <task id>")
		fmt.Println("Usage: task-cli update <task id> <task name>")
		fmt.Println("Usage: task-cli delete <task id>")
	default:
		fmt.Println("Usage: task-cli help")
		return
	}
}
