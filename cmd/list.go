package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available commands and their descriptions from your taskfile.yaml",
	Long:  `list available commands and their descriptions from your taskfile.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		// root flags
		if file, _ := cmd.Flags().GetString("task-file"); file != "" {
			tasks.FilePath = file
		}

		if err := tasks.Init(); err != nil {
			log.Fatal("task_list", err)
		}

		tasks.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
