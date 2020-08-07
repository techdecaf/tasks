package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/techdecaf/templates"
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

		cliVars, err  := cmd.Flags().GetStringToString("variable")
		if err != nil {
			log.Fatal("failed to set cli variables", err)
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

		for key, val := range cliVars {
			err := tasks.TemplateVars.Set(templates.Variable{
				Key: key,
				Value: val,
				OverrideEnv: true,
			})
			if err != nil {
				log.Fatal(fmt.Sprintf("could not set variable %s = %s", key, val), err)
			}
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
	var variables map[string]string
	
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringToStringVarP(&variables, "variable", "v" , nil, "overwrite environmental variables")
}
