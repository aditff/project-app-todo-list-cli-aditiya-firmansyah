package cmd

import (
	"fmt"
	"strconv"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("Invalid task number")
			return
		}
		err = todo.DeleteTask(index - 1)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task deleted.")
		}
	},
}
