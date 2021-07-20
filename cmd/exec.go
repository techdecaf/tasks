package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "execute a command using all resolved variables from your taskfile.yaml",
	Long:  `execute a command using all resolved variables from your taskfile.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("exec", "no arguments passed to exec.")
		}

		// root flags
		if file, _ := cmd.Flags().GetString("task-file"); file != "" {
			tasks.FilePath = file
		}

		cliVars, err := cmd.Flags().GetStringToString("variable")
		if err != nil {
			log.Fatal("failed to set cli variables", err)
		}
		SetEnvFrom(cliVars)

		if err := tasks.Init(); err != nil {
			log.Fatal("task_init", err)
		}

		// handle flags
		if log, _ := cmd.Flags().GetString("log"); log != "" {
			tasks.Options.LogLevel = (log == "true")
		}

		out, err := tasks.Execute(strings.Join(args, " "), "exec", "")
		if err != nil {
			log.Fatal("taskfie.execute", err)
		}

		fmt.Println(out)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	var variables map[string]string
	execCmd.Flags().StringToStringVarP(&variables, "variable", "v", nil, "overwrite environmental variables")

}
