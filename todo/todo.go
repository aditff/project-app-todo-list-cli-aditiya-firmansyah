package todo

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type Task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func NewTask(title string) Task {
	return Task{Title: title, Done: false}
}

func AddTask(task Task) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for _, t := range tasks {
		if strings.EqualFold(t.Title, task.Title) {
			return errors.New("task already exists")
		}
	}
	tasks = append(tasks, task)
	return saveTasks(tasks)
}

func ListTasks() error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "No\tTitle\tDone")
	for i, t := range tasks {
		doneStatus := "No"
		if t.Done {
			doneStatus = "Yes"
		}
		fmt.Fprintf(w, "%d\t%s\t%s\n", i+1, t.Title, doneStatus)
	}
	return w.Flush()
}

func MarkTaskDone(index int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if index < 0 || index >= len(tasks) {
		return errors.New("invalid task index")
	}
	tasks[index].Done = true
	return saveTasks(tasks)
}

func DeleteTask(index int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if index < 0 || index >= len(tasks) {
		return errors.New("invalid task index")
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	return saveTasks(tasks)
}

func SearchTasks(keyword string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	found := false
	for i, t := range tasks {
		if strings.Contains(strings.ToLower(t.Title), strings.ToLower(keyword)) {
			fmt.Printf("%d. %s [Done: %v]\n", i+1, t.Title, t.Done)
			found = true
		}
	}
	if !found {
		fmt.Println("No matching tasks found.")
	}
}
