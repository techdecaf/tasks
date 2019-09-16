package cmd

import (
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs a list of tasks as defined in your taskfile.yaml",
	Long:  `runs a list of tasks as defined in your taskfile.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		// root flags
		if file, _ := cmd.Flags().GetString("task-file"); file != "" {
			tasks.FilePath = file
		}

		if len(args) == 0 {
			args = append(args, "default")
		}

		if err := tasks.Init(); err != nil {
			log.Fatal("task_run", err)
		}

		// handle flags
		if silent, _ := cmd.Flags().GetBool("silent"); silent {
			tasks.Options.LogLevel = false
		}

		for _, task := range args {
			if err := tasks.Run(task); err != nil {
				log.Fatal(task, err)
			}
		}
		// fmt.Println("run called with %n commands", len(args))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
