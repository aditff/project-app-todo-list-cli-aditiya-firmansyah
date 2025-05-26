package todo

import (
	"encoding/json"
	"os"
)

const filePath = "data/todos.json"

// Load task list from JSON file
func loadTasks() ([]Task, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Task{}, nil
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

// Save task list to JSON file
func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
