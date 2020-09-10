package cmd

import (
	"fmt"
	"os"
	"strconv"

	"task_app/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do command description.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse arg:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invaild id ", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Println("Error in Deleting ", id, " :", err.Error())
				os.Exit(1)
			} else {
				fmt.Printf("Task %d: %s. is Marked Completed\n", id, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
