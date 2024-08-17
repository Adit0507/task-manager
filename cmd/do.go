package cmd

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			// Atoi = string to int
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task no.:", id)
				continue
			}

			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark %d as completed. Error: %s", id, err)
			} else {
				fmt.Printf("Marked %d as completed. ", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
