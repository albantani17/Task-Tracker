package task

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      Status `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateID(length int) string {
	id := make([]byte, length)
	for i := range id {
		id[i] = charset[rand.IntN(len(charset))]
	}
	return string(id)
}

func getColor(status Status) string {
	switch status {
	case Todo:
		return Red
	case InProgress:
		return Blue
	case Done:
		return Green
	default:
		return ""
	}
}

func AddTask(description string) {
	tasks := Load()
	task := Task{
		ID:          generateID(8),
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}
	fmt.Println("Task added successfully")
	tasks = append(tasks, task)
	Save(tasks)
}

func FindTasks() {
	tasks := Load()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for _, t := range tasks {
		color := getColor(t.Status)
		fmt.Printf("ID: %s\n", t.ID)
		fmt.Printf("Description: %s\n", t.Description)
		fmt.Printf("Status:"+color+" %s\n"+Reset, t.Status)
		fmt.Printf("Created At: %s\n", t.CreatedAt)
		fmt.Printf("Updated At: %s\n", t.UpdatedAt)
		fmt.Println()
	}
}

func FindTasksWithStatus(status Status) {
	tasks := Load()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	findTasks := []Task{}
	for _, t := range tasks {
		if t.Status != status {
			continue
		}
		findTasks = append(findTasks, t)
	}
	if len(findTasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for _, t := range findTasks {
		color := getColor(t.Status)
		fmt.Printf("ID: %s\n", t.ID)
		fmt.Printf("Description: %s\n", t.Description)
		fmt.Printf("Status:"+color+" %s\n"+Reset, t.Status)
		fmt.Printf("Created At: %s\n", t.CreatedAt)
		fmt.Printf("Updated At: %s\n", t.UpdatedAt)
		fmt.Println()
	}
}

func UpdateTask(id string, name string) {
	tasks := Load()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = name
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			Save(tasks)
			fmt.Println("Task updated successfully")
			return
		}
	}
}

func MarkTask(id string, status Status) {
	tasks := Load()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			Save(tasks)
			fmt.Println("Task marked successfully")
			return
		}
	}
}

func DeleteTask(id string) {
	tasks := Load()
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			Save(tasks)
			fmt.Println("Task deleted successfully")
			return
		}
	}
}
