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

		cliVars, err := cmd.Flags().GetStringToString("variable")
		if err != nil {
			logger.Fatal("failed to set cli variables", err)
		}
		SetEnvFrom(cliVars)

		if err := tasks.Init(); err != nil {
			logger.Fatal("task_list", err)
		}

		tasks.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	var variables map[string]string
	listCmd.Flags().StringToStringVarP(&variables, "variable", "v", nil, "overwrite environmental variables")

}
