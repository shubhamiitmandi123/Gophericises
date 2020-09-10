package cmd

import (
	"fmt"
	"os"
	"strings"

	"task_app/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Error adding task ", err.Error())
			os.Exit(1)
		}
		fmt.Println(task, " created! and Given id is :", id)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
