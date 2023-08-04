package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"meditasker/domain"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

var taskStrings = []string{
	"status: ",
	"project: ",
	"entered: ",
	"due: ",
	"uuid: ",
	"urgency: ",
}

const underln = "\033[4m"
const reset = "\033[0m"

// Here is a reference to your task manager. It could be initialized from somewhere else,
// for example from your main function.
var TaskManager *domain.TaskManager

var addTaskCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fullDescription := strings.Join(args, " ")
		task := parseTaskProperties(fullDescription)
		task.Entered = time.Now()

		err := TaskManager.AddTask(task)
		if err != nil {
			fmt.Println("Error adding task:", err)
			return
		}
		fmt.Println("Task added successfully")
	},
}

func underline(input string) string {
	return underln + input + reset
}

var getTasksCmd = &cobra.Command{
	Use:   "getTasks",
	Short: "Get all tasks",
	Long:  "This command gets all tasks from the task manager and prints them to the console.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := TaskManager.GetTasks()
		if err != nil {
			fmt.Println("Error getting tasks:", err)
			return
		}

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintln(w, underline("ID\tDescription\tStatus\tProject\tEntered\tDue\tUUID\tUrgency"))

		for _, task := range tasks {
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%f\n",
				task.ID,
				task.Description,
				task.Status,
				task.Project,
				task.Entered.Format(time.RFC3339),
				task.Due.Format(time.RFC3339),
				task.UUID,
				task.Urgency)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(addTaskCmd)
	rootCmd.AddCommand(getTasksCmd)
}

var rootCmd = &cobra.Command{
	Use:   "meditasker",
	Short: "meditasker is a command line task manager",
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
