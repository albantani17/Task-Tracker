package task

import (
	"encoding/json"
	"fmt"
	"os"
)

const storage = "storage/tasks.json"

func Load() []Task {
	var tasks []Task
	file, err := os.Open(storage)
	if err != nil {
		if err.Error() == "open tasks.json: The system cannot find the file specified." {
			os.Create(storage)
			return tasks
		}
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&tasks)
	return tasks
}

func Save(tasks []Task) {
	if err := os.MkdirAll("storage", os.ModePerm); err != nil {
		panic(err)
	}

	file, err := os.Create(storage)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(tasks)
	fmt.Println("Task saved successfully")
}
