package cmd

import (
	"github.com/spf13/cobra"
	"github.com/techdecaf/tasks/internal/taskfile"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Resolves all global variables and prints an export script",
	Long:  `Resolves all global variables and prints an export script`,
	Run: func(cmd *cobra.Command, args []string) {
		taskfile := &taskfile.TaskFile{}

		// root flags
		if file, _ := cmd.Flags().GetString("task-file"); file != "" {
			taskfile.FilePath = file
		}

		if err := taskfile.Init(); err != nil {
			log.Fatal("task_init", err)
		}
		taskfile.Export()
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
