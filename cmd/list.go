package cmd

import (
	"github.com/techdecaf/tasks/internal/taskfile"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available commands and their descriptions from your taskfile.yaml",
	Long:  `list available commands and their descriptions from your taskfile.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		taskfile := &taskfile.TaskFile{}

		// root flags
		if file, _ := cmd.Flags().GetString("task-file"); file != "" {
			taskfile.FilePath = file
		}

		if err := taskfile.Init(); err != nil {
			log.Fatal("task_init", err)
		}

		taskfile.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
