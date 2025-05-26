package cmd

import (
	"strings"
	"todo-cli/todo"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [keyword]",
	Short: "Search for tasks containing a keyword",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		keyword := strings.Join(args, " ")
		todo.SearchTasks(keyword)
	},
}
