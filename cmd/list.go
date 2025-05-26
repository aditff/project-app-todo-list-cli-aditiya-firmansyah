package cmd

import (
	"fmt"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		err := todo.ListTasks()
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}
