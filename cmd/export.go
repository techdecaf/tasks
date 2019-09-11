package cmd

import (
	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Resolves all global variables and prints an export script",
	Long:  `Resolves all global variables and prints an export script`,
	Run: func(cmd *cobra.Command, args []string) {
		// root flags
		if file, _ := cmd.Flags().GetString("task-file"); file != "" {
			tasks.FilePath = file
		}

		if err := tasks.Init(); err != nil {
			log.Fatal("task_init", err)
		}
		tasks.Export()
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
