package cmd

import (
	"fmt"
	"strings"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task title]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		if strings.TrimSpace(title) == "" {
			fmt.Println("Task title cannot be empty.")
			return
		}

		task := todo.NewTask(title)
		err := todo.AddTask(task)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task added:", title)
		}
	},
}
