package cmd

import (
	"fmt"
	"os"

	"task_app/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List command Description.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Somethig went wrong ", err.Error())
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("you have no tasks, Why not take a vecation?")
		} else {
			fmt.Println("You have Following Tasks to Do")
			for i, v := range tasks {
				fmt.Printf("%d : %s\n", i+1, v.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
